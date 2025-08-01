package tgbot

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (b *Bot) Run(ctx context.Context) {
	logrus.Infof("Starting bot %s", b.botAPI.Self.UserName)

	for {
		select {
		case <-ctx.Done():
			return
		case <-b.done:
			return
		case u, ok := <-b.updatesChan:
			if !ok {
				return
			}

			if u.Message == nil {
				continue
			}

			logrus.Infof("[%s] %s", u.Message.From.UserName, u.Message.Text)

			if !u.Message.IsCommand() {
				continue
			}

			f, ok := b.commands[u.Message.Command()]
			if !ok {
				b.sendText(u, fmt.Sprintf("unknown command: %s", u.Message.Command()))
				continue
			}

			msg := f(ctx, u)

			b.sendMsg(u, msg)
		}
	}
}
