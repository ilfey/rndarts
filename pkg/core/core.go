package core

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Core struct {
	Bot *tgbotapi.BotAPI
	*logrus.Logger
}

func NewCore(bot *tgbotapi.BotAPI, logger *logrus.Logger) *Core {
	return &Core{
		Bot:    bot,
		Logger: logger,
	}
}
