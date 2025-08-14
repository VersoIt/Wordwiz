package main

import (
	"context"
	"fmt"
	trmngr "github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wordwiz/config"
	"wordwiz/internal/domain/service/user"
	smartgenuc "wordwiz/internal/domain/usecase/music_smart_generator"
	"wordwiz/internal/infrastructure/client/gemini"
	postgres "wordwiz/internal/infrastructure/pg"
	"wordwiz/internal/infrastructure/pg/migrator"
	userrepo "wordwiz/internal/infrastructure/repository/pg/user_repo"
	"wordwiz/internal/transport/tgbot"
	"wordwiz/internal/worker"
)

func main() {
	cfg := config.MustGet()

	pg, err := postgres.New(cfg)
	if err != nil {
		logrus.Panic(err)
	}

	pgMigrator := migrator.New(pg)

	err = pgMigrator.Migrate()
	if err != nil {
		logrus.Panic(err)
	}

	txManager := manager.Must(trmngr.NewDefaultFactory(pg))

	userRepo := userrepo.New(pg, trmngr.DefaultCtxGetter)
	userService := user.New(userRepo)
	aiClient := gemini.New(cfg, http.Client{})

	musicGeneratorUC := smartgenuc.New(
		cfg,
		userRepo,
		txManager,
		userService,
		aiClient,
	)

	bot, err := tgbot.New(cfg)
	if err != nil {
		logrus.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*cfg.ServerCloseTimeoutMS)
	defer cancel()

	bot.HandleCommand("start", func(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "Hi! Send me some words and I'll generate lyrics for you based on them.\nExample: /gen sun, beach, sea")
	})

	bot.HandleCommand("help", func(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig {
		return tgbotapi.NewMessage(update.Message.Chat.ID, "/gen - generate song lyrics from given words")
	})

	bot.HandleCommand("gen", func(ctx context.Context, update tgbotapi.Update) tgbotapi.MessageConfig {

		stats, verses, err := musicGeneratorUC.GenerateWithStats(ctx, update.Message.CommandArguments(), update.Message.From.ID)
		if err != nil {
			logrus.Errorf("[%d]: %v", update.Message.From.ID, err)

			return tgbotapi.NewMessage(
				update.Message.Chat.ID,
				fmt.Sprintf("Error while generating song lyrics: %v", err),
			)
		}

		raw := verses.String()

		return tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf(
				"%s\n\nRequests per month: %d out of %d",
				raw,
				stats.GenerationsPerMonth,
				cfg.MaxGenerationsPerMonth,
			),
		)
	})

	go func() {
		bot.Run(ctx)
	}()

	cronWorker := worker.New(userRepo)

	go func() {
		err := cronWorker.Run(ctx)
		if err != nil {
			logrus.Panicf("Worker error: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt)
	<-done

	cancel()
}
