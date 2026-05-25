// Package config loads runtime configuration from environment variables.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	Env                  string
	GRPCAddr             string
	HealthAddr           string
	HealthEnable         bool
	GRPCReflectionEnable bool
}

func Load() (Config, error) {
	cfg := Config{}

	cfg.Env = strings.TrimSpace(os.Getenv("APP_ENV"))
	if cfg.Env == "" {
		cfg.Env = "local"
	}

	// Dotenv is ONLY for local dev; containers/prod should inject env.
	if cfg.Env == "local" {
		if envFile := strings.TrimSpace(os.Getenv("ENV_FILE")); envFile != "" {
			_ = godotenv.Load(envFile)
		} else if _, err := os.Stat(filepath.Clean(".env")); err == nil {
			_ = godotenv.Load()
		}
	}

	cfg.GRPCAddr = strings.TrimSpace(os.Getenv("GRPC_ADDR"))
	if cfg.GRPCAddr == "" {
		cfg.GRPCAddr = ":50051"
	}

	cfg.HealthAddr = strings.TrimSpace(os.Getenv("HEALTH_ADDR"))
	if cfg.HealthAddr == "" {
		cfg.HealthAddr = ":8080"
	}

	cfg.HealthEnable = true
	if v := strings.TrimSpace(os.Getenv("HEALTH_ENABLED")); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return Config{}, fmt.Errorf("invalid HEALTH_ENABLED: %w", err)
		}
		cfg.HealthEnable = b
	}

	cfg.GRPCReflectionEnable = false
	if v := strings.TrimSpace(os.Getenv("GRPC_REFLECTION_ENABLED")); v != "" {
		b, err := strconv.ParseBool(v)
		if err != nil {
			return Config{}, fmt.Errorf("invalid GRPC_REFLECTION_ENABLED: %w", err)
		}
		cfg.GRPCReflectionEnable = b
	}

	return cfg, nil
}
