package global

import (
	"github.com/spf13/viper"
	"log"
	"sync"
)

var instance *Config
var once sync.Once

type Config struct {
	jwtKey string
}

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

	cfg.jwtKey = viper.GetString("jwt.key")
}

func (c *Config) SetJwtKey() string {
	return c.jwtKey
}
