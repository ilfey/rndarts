package app

import (
	"context"
	"time"

	"main/pkg/app/controller"
	"main/pkg/cmd"
	"main/pkg/core"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	*core.Core
	*controller.Controller
}

func NewApp(core *core.Core) *App {
	return &App{
		Core:       core,
		Controller: controller.NewController(),
	}
}

func (a *App) onCommandsRegister() {
	a.Controller.Init()

	cmds := a.Controller.ToBotCommands()

	a.Bot.Request(tgbotapi.NewSetMyCommands(cmds...))

}

func (a *App) onStart() {
	a.Logger.Infof("Authorized on account %s", a.Bot.Self.UserName)

	a.Logger.Info("Registering commands")
	a.onCommandsRegister()
}

func (a *App) onUpdate(update tgbotapi.Update) {
	// If the message is a command
	if update.Message != nil && update.Message.IsCommand() {
		for _, c := range a.Commands {
			if c.Name() == update.Message.Command() {
				a.Logger.Infof("%s call command: %s", update.Message.From.UserName, update.Message.Text)

				// Create context timeout
				parent, cancel := context.WithTimeout(context.Background(), 3*time.Second)

				// Create command context
				ctx := cmd.NewContext(
					parent,
					&update,
					a.Core,
				)

				// Handle command
				c.Handle(ctx)

				// Cancel context timeout
				cancel()

				// Stop the loop
				break
			}
		}
	}
}

func (a *App) Serve() {
	a.onStart()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := a.Bot.GetUpdatesChan(u)

	for update := range updates {
		a.onUpdate(update)
	}
}
