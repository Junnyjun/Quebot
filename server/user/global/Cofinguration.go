package global

import (
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var instance *Config
var once sync.Once

type Config struct {
	jwtKey    string
	expiresAt string
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
	}

	viper.SetConfigName("config-" + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Println("[Configuration] Config file loaded : ", viper.ConfigFileUsed())
	cfg.jwtKey = viper.GetString("jwt.key")
	cfg.expiresAt = viper.GetString("jwt.expiresAt")
}

func (c *Config) JwtKey() string {
	return c.jwtKey
}

func (c *Config) ExpiresAt() time.Time {
	durationStr := c.expiresAt + "s"
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		log.Fatalf("Invalid duration format: %s", err)
		return time.Time{}
	}
	return time.Now().Add(duration).UTC()
}
