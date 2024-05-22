package main

import (
	"main/pkg/app"
	"main/pkg/config"
	"main/pkg/core"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	ConfigureLogger()

	if err := config.ReadConfig(); err != nil {
		panic(err)
	}

	if err := config.ValidateConfig(); err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(config.GetBotToken())
	if err != nil {
		panic(err)
	}

	c := core.NewCore(bot, logrus.StandardLogger())

	a := app.NewApp(c)

	a.Serve()
}

func ConfigureLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false,
		ForceColors:   true,
		ForceQuote:    true,
	})
}
