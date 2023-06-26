package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v8"
)

type config struct {
	Port int `env:"PORT" envDefault:"3000"`
}

type Config interface {
	GetPort() string
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
