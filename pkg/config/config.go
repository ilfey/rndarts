package config

import "os"

type Config struct {
	BotToken string
}

func NewConfig() *Config {
	return &Config{
		BotToken: getEnv("BOT_TOKEN", ""),
	}
}

// Get environment variable
func getEnv(key, fallback string) string {
	// If value exists, return it
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		return fallback
	}
}
