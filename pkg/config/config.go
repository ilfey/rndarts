package config

import (
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type key = string

const (
	BotTokenKey       key = "bot_token"
	AdminsKey         key = "admins" // TODO: not used
	ChannelsKey       key = "channels"
	NotifyChannelsKey key = "notify_channels"
)

var (
	ErrBotTokenNotSet       = errors.New("bot_token not set")
	ErrAdminsNotSet         = errors.New("channels not set")
	ErrChannelsNotSet       = errors.New("channels not set")
	ErrNotifyChannelsNotSet = errors.New("notify_channels not set")
)

func ReadConfig() error {
	if viper.IsSet("config") && viper.GetString("config") != "" {
		logrus.Debug("Setting config file from flag")

		viper.SetConfigFile(viper.GetString("config"))
	} else {
		logrus.Debug("Setting default config name and type")

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")

		logrus.Debug("Adding default config paths")

		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.rndanart")
	}

	logrus.Info("Reading config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func SaveConfig() error {
	logrus.Info("Saving config")

	return viper.WriteConfig()
}

func GetBotToken() string {
	logrus.Debug("Getting bot token")

	return viper.GetString(BotTokenKey)
}

func SetBotToken(token string) {
	logrus.Debug("Setting bot token")

	viper.Set(BotTokenKey, token)
}

func GetAdmins() []string {
	logrus.Debug("Getting admins")

	return viper.GetStringSlice(AdminsKey)
}

func SetAdmins(admins []string) {
	logrus.Debug("Setting admins")

	viper.Set(AdminsKey, admins)
}

func GetChannels() []int64 {
	logrus.Debug("Getting channels")
	slice := viper.GetStringSlice(ChannelsKey)

	channels := make([]int64, len(slice)-1)

	for _, channel := range viper.GetStringSlice(ChannelsKey) {
		i, err := strconv.ParseInt(channel, 10, 64)
		if err != nil {
			logrus.Panicf("Failed to parse channel: %s", err)
		}

		channels = append(channels, i)
	}

	return channels
}

func SetChannels(channels []string) {
	logrus.Debug("Setting channels")

	viper.Set(ChannelsKey, channels)
}

func GetNotifyChannels() string {
	logrus.Debug("Getting notify channels")

	return viper.GetString(NotifyChannelsKey)
}

func SetNotifyChannels(channels string) {
	logrus.Debug("Setting notify channels")

	viper.Set(NotifyChannelsKey, channels)
}

func ValidateConfig() error {
	if !viper.IsSet(BotTokenKey) {
		return ErrBotTokenNotSet
	}

	if !viper.IsSet(ChannelsKey) {
		return ErrChannelsNotSet
	}

	if !viper.IsSet(NotifyChannelsKey) {
		return ErrNotifyChannelsNotSet
	}

	return nil
}
