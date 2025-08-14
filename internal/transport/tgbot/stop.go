package tgbot

func (b *Bot) Stop() {
	close(b.done)
}
