package config

import "os"

const (
	DefaultPort string = "8080"
)

type Config struct {
	Port string
}

func New() *Config {
	port := DefaultPort
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	return &Config{
		Port: port,
	}
}
