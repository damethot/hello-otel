package internal

import (
	"github.com/caarlos0/env/v10"
)

type Config struct {
	PongerAddr  string `env:"PONGER_ADDR"`
	ServiceName string `env:"SERVICE_NAME"`
}

func ParseEnv() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil

}
