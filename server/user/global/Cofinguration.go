package global

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

var instance *Config
var once sync.Once

type Config struct {
	jwtKey string
}

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
		loadConfig(instance)
	})
	return instance
}

func loadConfig(cfg *Config) {
	environment := os.Getenv("ENV")
	if environment == "" {
		log.Panic("ENV is not set")
		panic(`ENV is not set`)
	}

	viper.SetConfigName("config-" + environment)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	cfg.jwtKey = viper.GetString("jwt.key")
}

func (c *Config) JwtKey() string {
	return c.jwtKey
}
