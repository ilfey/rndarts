package config

import (
	"errors"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type key = string

const (
	BotTokenKey key = "BOT_TOKEN"
	AdminsKey   key = "ADMINS"
	ChannelsKey key = "CHANNELS"
)

func ReadConfig() error {

	logrus.Debug("Setting config name and type")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	logrus.Debug("Adding config paths")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.rndanart")

	logrus.Info("Reading config")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("Error reading config file: %s", err)
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

func GetChannels() []string {
	logrus.Debug("Getting channels")

	return viper.GetStringSlice(ChannelsKey)
}

func SetChannels(channels []string) {
	logrus.Debug("Setting channels")

	viper.Set(ChannelsKey, channels)
}

func ValidateConfig() error {
	if !viper.IsSet(BotTokenKey) {
		return errors.New("bot token not set")
	}

	return nil
}
