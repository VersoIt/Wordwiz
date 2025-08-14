package tgbot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"wordwiz/config"
)

type CommandHandler func(
	ctx context.Context,
	update tgbotapi.Update,
) tgbotapi.MessageConfig

type Bot struct {
	commands    map[string]CommandHandler
	updatesChan tgbotapi.UpdatesChannel
	botAPI      *tgbotapi.BotAPI
	done        chan struct{}
}

func New(cfg config.Config) (*Bot, error) {
	botAPI, err := tgbotapi.NewBotAPI(cfg.TGBot.Token)
	if err != nil {
		return nil, err
	}

	botAPI.Debug = cfg.TGBot.Debug

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = int(cfg.TGBot.TimeoutMS)

	updatesChan, err := botAPI.GetUpdatesChan(updateConfig)
	if err != nil {
		return nil, err
	}

	return &Bot{
		commands:    make(map[string]CommandHandler),
		updatesChan: updatesChan,
		botAPI:      botAPI,
		done:        make(chan struct{}),
	}, nil
}
