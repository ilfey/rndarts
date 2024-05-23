package tasks

import (
	"main/pkg/api/nekos"
	"main/pkg/config"
	"main/pkg/core"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Send(c *core.Core, getKind func() nekos.Kind) {
	bot := c.Bot
	logger := c.Logger

	kind := getKind()

	logger.Debugf("Get random image for kind: %s", kind)
	response, err := nekos.GetRandomImage(kind)

	if err != nil {
		logger.Errorf("Get random image error: %s", err)
		return
	}

	for _, channel := range config.GetChannels() {
		logger.Debugf("Send image to channel: %d", channel)

		doc := tgbotapi.NewPhoto(channel, tgbotapi.FileURL(*response.Results[0].URL))
		// doc.Caption = *response.Results[0].SourceURL

		bot.Send(doc)
	}
}
