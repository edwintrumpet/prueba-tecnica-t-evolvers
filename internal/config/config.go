package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v8"
)

type config struct {
	Port int `env:"PORT" envDefault:"3000"`
	DB   struct {
		Host     string `env:"DB_HOST" envDefault:"localhost"`
		User     string `env:"DB_USER" envDefault:"postgres"`
		Password string `env:"DB_PASSWORD" envDefault:"postgres"`
		DBName   string `env:"DB_NAME"`
		Port     int    `env:"DB_PORT" envDefault:"5432"`
	}
	Redis struct {
		Addr string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	}
	Env string `env:"ENV" envDefault:"prod"`
}

type Config interface {
	GetPort() string
	DBDns() string
	IsDev() bool
	RedisAddr() string
}

// New creates and populates a config object with all the environment variables
// and returns a Config interface to provide the necessary environment data
func New() Config {
	var c config

	if err := env.Parse(&c); err != nil {
		log.Fatalf("error parsing env variables %s", err.Error())
	}

	return &c
}

func (c *config) GetPort() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *config) DBDns() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.DB.Host,
		c.DB.User,
		c.DB.Password,
		c.DB.DBName,
		c.DB.Port,
	)
}

func (c *config) IsDev() bool {
	return c.Env == "dev"
}

func (c *config) RedisAddr() string {
	return c.Redis.Addr
}
