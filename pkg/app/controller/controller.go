package controller

import (
	"fmt"

	"main/pkg/api/nekos"
	"main/pkg/app/handlers"
	"main/pkg/cmd"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	Commands []cmd.Command
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Init() {
	c.Commands = []cmd.Command{
		cmd.NewCommand(
			&cmd.Options{
				Name: "start",
				Help: "Начать работу",
			},
			handlers.Start,
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "waifu",
				Help: "Получить случайную waifu",
			},
			handlers.NewNekosImageHandler(nekos.WAIFU),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "kitsu",
				Help: "Получить случайную kitsu",
			},
			handlers.NewNekosImageHandler(nekos.KITSU),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "neko",
				Help: "Получить случайную neko",
			},
			handlers.NewNekosImageHandler(nekos.NEKO),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "husbando",
				Help: "Получить случайного husbando",
			},
			handlers.NewNekosImageHandler(nekos.HUSBANDO),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "help",
				Help: "Показать справку",
			},
			c.helpHandler,
		),
	}
}

func (c *Controller) ToBotCommands() []tgbotapi.BotCommand {
	commands := make([]tgbotapi.BotCommand, 0, len(c.Commands))

	for _, cmd := range c.Commands {
		commands = append(commands, *cmd.ToBotCommand())
	}

	return commands
}
func (c *Controller) helpHandler(ctx cmd.Context) {
	helpText := "*Help:*\n"

	for _, cmd := range c.Commands {
		helpText += fmt.Sprintf("/%s - %s\n", cmd.Name(), cmd.Help())
	}

	msg := tgbotapi.NewMessage(ctx.GetChatID(), helpText)
	msg.ParseMode = tgbotapi.ModeMarkdown
	msg.ReplyToMessageID = ctx.GetMessageID()

	ctx.Send(msg)
}
