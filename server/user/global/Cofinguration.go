package global

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

type Config struct {
	ServerAddress string
	LogLevel      string
}

var instance *Config
var once sync.Once

func GetInstance(environment string) *Config {
	once.Do(func() {
		instance = &Config{}
		loadConfig(instance, environment)
	})
	return instance
}

func loadConfig(cfg *Config, environment string) {
	viper.SetConfigName("config-" + environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	cfg.ServerAddress = viper.GetString("server.address")
	cfg.LogLevel = viper.GetString("log.level")
}
