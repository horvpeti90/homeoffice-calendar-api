package config

import (
	"os"

	"github.mpi-internal.com/hu/homeoffice-calendar-api/internal/transport"
)

type LoggerConfig struct {
	LogLevel string
}

type Config struct {
	Logger    LoggerConfig
	Transport transport.Config
}

func New() Config {
	return Config{
		Logger: LoggerConfig{
			LogLevel: getenv("LOG_LEVEL", "debug"),
		},
		Transport: transport.Config{
			Address: getenv("TRANSPORT_ADDRESS", ":8080"),
		},
	}
}

func getenv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}
