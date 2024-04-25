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

func GetInstance(env string) *Config {
	once.Do(func() {
		instance = &Config{}
		loadConfig(instance, env)
	})
	return instance
}

func loadConfig(cfg *Config, env string) {
	log.Println("[Configuration] ENV : ", env)
	if env == "" {
		log.Panic("ENV is not set")
		panic(`ENV is not set`)
	}

	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Println("[Configuration] Config file loaded : ", viper.ConfigFileUsed())
	cfg.jwtKey = viper.GetString("jwt.key")
}

func (c *Config) JwtKey() string {
	return c.jwtKey
}
