package apiserver

import (
	"github.com/KebaCorp/TechnologyStackAPI/internal/app/store"
)

// Config ...
type Config struct {
	BindAddr   string `toml:"bind_addr"`
	LogLevel   string `toml:"log_level"`
	CorsOrigin string `toml:"cors_origin"`
	Store      *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		LogLevel:   "debug",
		Store:      store.NewConfig(),
		CorsOrigin: "http://localhost:7082",
	}
}
