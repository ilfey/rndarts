package handlers

import (
	"main/pkg/api/nekos"
	"main/pkg/cmd"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(ctx cmd.Context) {
	u := ctx.GetUpdate()

	msg := tgbotapi.NewMessage(u.Message.Chat.ID, "Привет, это бот для рандомных аниме картиночек!")
	msg.ReplyToMessageID = u.Message.MessageID

	ctx.Send(msg)
}

func NewNekosImageHandler(kind nekos.Kind) cmd.Handler {
	return func(ctx cmd.Context) {
		logger := ctx.GetLogger()

		logger.Infof("Get random image for kind: %s", kind)
		response, err := nekos.GetRandomImage(kind)

		if err != nil {
			logger.Errorf("Get random image error: %s", err)
			return
		}

		doc := tgbotapi.NewPhoto(ctx.GetChatID(), tgbotapi.FileURL(*response.Results[0].URL))
		doc.Caption = *response.Results[0].SourceURL
		doc.ParseMode = tgbotapi.ModeMarkdown
		doc.ReplyToMessageID = ctx.GetMessageID()

		ctx.Send(doc)
	}
}
