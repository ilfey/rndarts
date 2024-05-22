package cmd

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Handler = func(Context)

type Command interface {
	Name() string
	Help() string
	Handle(Context)
	ToBotCommand() *tgbotapi.BotCommand
}

type Options struct {
	Name string
	Help string
}

type CommandImpl struct {
	options *Options
	handler func(ctx Context)
}

func (c *CommandImpl) Name() string {
	return c.options.Name
}

func (c *CommandImpl) Help() string {
	return c.options.Help
}

func (c *CommandImpl) ToBotCommand() *tgbotapi.BotCommand {
	return &tgbotapi.BotCommand{
		Command:     c.options.Name,
		Description: c.options.Help,
	}
}

func (c *CommandImpl) Handle(ctx Context) {
	c.handler(ctx)
}

func NewCommand(op *Options, handler Handler) Command {
	return &CommandImpl{
		options: op,
		handler: handler,
	}
}
