package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

func (b *Bot) sendMsg(
	update tgbotapi.Update,
	msg tgbotapi.MessageConfig,
) {
	_, err := b.botAPI.Send(msg)
	if err != nil {
		logrus.Errorf("[%s] %s", update.Message.From.UserName, err.Error())
	}
}
