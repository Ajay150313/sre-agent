package config

import (
	"os"
)

type Config struct {
	OpenAIKey      string
	PrometheusURL  string
	Port           string
	LogLevel       string
	MaxRetries     int
}

func Load() *Config {
	return &Config{
		OpenAIKey:      getEnv("OPENAI_API_KEY", ""),
		PrometheusURL:  getEnv("PROMETHEUS_URL", "http://localhost:9090"),
		Port:           getEnv("PORT", "8080"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		MaxRetries:     3,
	}
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
