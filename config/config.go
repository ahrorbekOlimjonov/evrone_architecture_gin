package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	// Config -.
	Config struct {
		App  App
		HTTP HTTP
		Log  Log
		PG   PG
		RMQ  RMQ
	}

	// App -.
	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	// Log -.
	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	// RMQ -.
	RMQ struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	fmt.Println("config: ", cfg)

	return cfg, nil
}
