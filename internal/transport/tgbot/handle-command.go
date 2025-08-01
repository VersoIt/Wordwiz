package tgbot

func (b *Bot) HandleCommand(
	command string,
	f CommandHandler,
) {
	b.commands[command] = f
}
