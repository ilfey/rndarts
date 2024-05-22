package cmd

import (
	"context"
	"main/pkg/core"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type CommandContext struct {
	parent context.Context

	*tgbotapi.Update
	*core.Core
}

func NewContext(parent context.Context, update *tgbotapi.Update, core *core.Core) *CommandContext {
	return &CommandContext{
		parent: parent,
		Update: update,
		Core:   core,
	}
}

type Context interface {
	GetUpdate() *tgbotapi.Update

	GetCore() *core.Core
	GetBot() *tgbotapi.BotAPI
	GetLogger() *logrus.Logger

	GetChatID() int64
	GetUserID() int64
	GetUserName() string
	GetMessageID() int
	GetText() string
	Send(msg tgbotapi.Chattable)

}

func (c *CommandContext) GetUpdate() *tgbotapi.Update {
	return c.Update
}

func (c *CommandContext) GetCore() *core.Core {
	return c.Core
}

func (c *CommandContext) GetBot() *tgbotapi.BotAPI {
	return c.Core.Bot
}

func (c *CommandContext) GetLogger() *logrus.Logger {
	return c.Core.Logger
}

// UTILS

func (c *CommandContext) GetChatID() int64 {
	return c.GetUpdate().Message.Chat.ID
}

func (c *CommandContext) GetUserID() int64 {
	return c.GetUpdate().Message.From.ID
}

func (c *CommandContext) GetUserName() string {
	return c.GetUpdate().Message.From.UserName
}

func (c *CommandContext) GetMessageID() int {
	return c.GetUpdate().Message.MessageID
}

func (c *CommandContext) GetText() string {
	return c.GetUpdate().Message.Text
}

func (c *CommandContext) Send(msg tgbotapi.Chattable) {
	c.GetBot().Send(msg)
}

// BASE CONTEXT

func (c *CommandContext) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

func (c *CommandContext) Done() <-chan struct{} {
	return c.parent.Done()
}

func (c *CommandContext) Err() error {
	return c.parent.Err()
}

func (c *CommandContext) Value(key any) any {
	return c.parent.Value(key)
}
