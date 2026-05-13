package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Port string
	Env  string

	AuthServiceURL string
	ChatServiceURL string
	TourServiceURL string

	JWTSecret string
	JWTTTL    time.Duration

	HTTPTimeout time.Duration

	LogLevel string
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("ENV", "local"),

		AuthServiceURL: getEnv("AUTH_SERVICE_URL", ""),
		ChatServiceURL: getEnv("CHAT_SERVICE_URL", ""),
		TourServiceURL: getEnv("TOUR_SERVICE_URL", ""),

		JWTSecret: getEnv("JWT_SECRET", ""),
		LogLevel:  getEnv("LOG_LEVEL", "debug"),
	}

	if cfg.AuthServiceURL == "" {
		return nil, fmt.Errorf("AUTH_SERVICE_URL is required")
	}

	if cfg.ChatServiceURL == "" {
		return nil, fmt.Errorf("CHAT_SERVICE_URL is required")
	}

	if cfg.TourServiceURL == "" {
		return nil, fmt.Errorf("TOUR_SERVICE_URL is required")
	}

	if cfg.JWTSecret == "" {
		return nil, fmt.Errorf("JWT_SECRET is required")
	}

	jwtTTL := getEnv("JWT_TTL", "24h")
	parsedTTL, err := time.ParseDuration(jwtTTL)
	if err != nil {
		return nil, fmt.Errorf("invalid JWT_TTL: %w", err)
	}
	cfg.JWTTTL = parsedTTL

	timeoutMs := getEnv("HTTP_TIMEOUT_MS", "5000")
	ms, err := strconv.Atoi(timeoutMs)
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP_TIMEOUT_MS: %w", err)
	}
	cfg.HTTPTimeout = time.Duration(ms) * time.Millisecond

	return cfg, nil
}

func getEnv(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}