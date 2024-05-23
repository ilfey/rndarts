package main

import (
	"main/pkg/app"
	"main/pkg/config"
	"main/pkg/core"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// Setup flags
	pflag.StringP("config", "c", "", "config file path")
	pflag.BoolP("debug", "d", false, "debug mode")
}

func main() {
	// Parse flags
	pflag.Parse()

	// Bind flags to config
	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		logrus.Panicf("Failed to bind flags: %s", err)
	}

	// Set debug mode
	if viper.GetBool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	ConfigureLogger()

	// Read config file
	if err := config.ReadConfig(); err != nil {
		switch err {
		case config.ErrBotTokenNotSet:
			logrus.Panicf("Not found config file: %s", err)
		default:
			logrus.Panicf("Failed to read config file: %s", err)
		}
	}

	if err := config.ValidateConfig(); err != nil {
		switch err {
		case config.ErrBotTokenNotSet:
			logrus.Panicf("BOT_TOKEN not set: %s", err)
		default:
			logrus.Panicf("Failed to validate config: %s", err)
		}
	}

	logrus.Debug("Connecting bot")
	bot, err := tgbotapi.NewBotAPI(config.GetBotToken())
	if err != nil {
		logrus.Panicf("Failed connecting: %s", err)
	}

	logrus.Debug("Creating core")
	c := core.NewCore(bot, logrus.StandardLogger())

	logrus.Debug("Creating app")
	a := app.NewApp(c)

	logrus.Debug("Start app serving")
	a.Serve()
}

func ConfigureLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: false,
		ForceColors:   true,
		ForceQuote:    true,
	})
}
