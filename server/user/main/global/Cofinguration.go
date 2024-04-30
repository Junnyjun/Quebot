package global

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
	"time"
)

var instance *Config
var once sync.Once

type Config struct {
	jwtKey    string
	expiresAt string

	port string
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

	err := godotenv.Load(".env." + env)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	cfg.jwtKey = os.Getenv("JWT_KEY")
	log.Println("[Configuration] JWT_KEY : " + cfg.JwtKey())

	cfg.expiresAt = os.Getenv("EXPIRES_AT")
	log.Println("[Configuration] EXPIRES_AT : " + cfg.expiresAt)

	cfg.port = os.Getenv("PORT")
	log.Println("[Configuration] PORT : " + cfg.Port())
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

func (c *Config) Port() string {
	return "127.0.0.1:" + c.port
}
