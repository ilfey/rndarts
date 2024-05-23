package app

import (
	"context"
	"time"

	"main/pkg/app/controller"
	"main/pkg/cmd"
	"main/pkg/core"
	"main/pkg/worker"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type App struct {
	*core.Core
	Controller *controller.Controller
	Worker     *worker.Worker
}

func NewApp(core *core.Core) *App {
	return &App{
		Core:       core,
		Controller: controller.NewController(core),
		Worker:     worker.NewWorker(core),
	}
}

func (a *App) onStart() {
	a.Logger.Infof("[onStart] Authorized on account %s", a.Bot.Self.UserName)

	a.Logger.Info("[onStart] Registering commands")
	a.onCommandsRegister()

	a.Logger.Info("[onStart] Starting worker")
	a.onStartWorker()
}

func (a *App) onCommandsRegister() {
	a.Logger.Info("[onCommandsRegister] Loading commands")
	a.Controller.LoadCommands()

	cmds := a.Controller.ToBotCommands()

	a.Bot.Request(tgbotapi.NewSetMyCommands(cmds...))
}

func (a *App) onStartWorker() {
	a.Logger.Info("[onStartWorker] Loading tasks")
	a.Controller.LoadTasks()

	a.Logger.Info("[onStartWorker] Adding tasks")
	a.Worker.AddTask(a.Controller.Tasks...)

	go a.Worker.Start()
}

func (a *App) onMessage(update *tgbotapi.Update) {
	if update.Message.IsCommand() {
		for _, c := range a.Controller.Commands {
			if c.Name() == update.Message.Command() {
				a.Logger.Infof("[onMessage] %s call command: %s", update.Message.From.UserName, update.Message.Text)

				// Create context timeout
				parent, cancel := context.WithTimeout(context.Background(), 3*time.Second)

				// Create command context
				ctx := cmd.NewContext(
					parent,
					update,
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

func (a *App) onChannelPost(update *tgbotapi.Update) {
	a.Logger.Infof("[onChannelPost] New channel post: %s", update.ChannelPost.Text)
}

func (a *App) onStop() {
	a.Logger.Info("[onStop] Stopping worker")
	a.Worker.Stop()
}

func (a *App) onUpdate(update tgbotapi.Update) {
	// If update is message
	if update.Message != nil {
		a.Logger.Info("[onUpdate] New message")
		a.onMessage(&update)
	}

	// If update is channel post
	if update.ChannelPost != nil {
		a.Logger.Info("[onUpdate] New channel post")
		a.onChannelPost(&update)
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

	a.onStop()
}
