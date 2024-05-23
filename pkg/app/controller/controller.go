package controller

import (
	"fmt"

	"main/pkg/api/nekos"
	"main/pkg/app/handlers"
	"main/pkg/app/tasks"
	"main/pkg/cmd"
	"main/pkg/config"
	"main/pkg/core"
	"main/pkg/worker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Controller struct {
	core     *core.Core
	Commands []cmd.Command
	Tasks    []*worker.Task
}

func NewController(core *core.Core) *Controller {
	return &Controller{
		core: core,
	}
}

// Load commands
func (c *Controller) LoadCommands() {
	c.Commands = []cmd.Command{
		cmd.NewCommand(
			&cmd.Options{
				Name: "start",
				Help: "Начать работу",

				IsVisible: true,
			},
			handlers.Start,
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "waifu",
				Help: "Получить случайную вайфу",

				IsVisible: true,
			},
			handlers.NewNekosImageHandler(nekos.WAIFU),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "kitsu",
				Help: "Получить случайную лисичку",

				IsVisible: true,
			},
			handlers.NewNekosImageHandler(nekos.KITSU),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "neko",
				Help: "Получить случайную неко-тян",

				IsVisible: true,
			},
			handlers.NewNekosImageHandler(nekos.NEKO),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "husbando",
				Help: "Получить случайного хусбандо",

				IsVisible: true,
			},
			handlers.NewNekosImageHandler(nekos.HUSBANDO),
		),
		cmd.NewCommand(
			&cmd.Options{
				Name: "help",
				Help: "Показать справку",

				IsVisible: true,
			},
			c.helpHandler,
		),
	}
}

// Load tasks
func (c *Controller) LoadTasks() {
	var kind nekos.Kind

	c.Tasks = []*worker.Task{
		worker.NewTask(
			config.GetNotifyChannels(),
			func() {
				tasks.Send(c.core, func() nekos.Kind {
					switch kind {
					case nekos.WAIFU:
						kind = nekos.KITSU
					case nekos.KITSU:
						kind = nekos.NEKO
					case nekos.NEKO:
						kind = nekos.HUSBANDO
					case nekos.HUSBANDO:
						kind = nekos.WAIFU
					default:
						kind = nekos.WAIFU
					}

					return kind
				})
			},
		),
	}
}

// Convert commands to bot commands
func (c *Controller) ToBotCommands() []tgbotapi.BotCommand {
	commands := make([]tgbotapi.BotCommand, 0, len(c.Commands))

	for _, cmd := range c.Commands {
		if !cmd.IsVisible() {
			continue
		}

		commands = append(commands, *cmd.ToBotCommand())
	}

	return commands
}

// Help hendler
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
