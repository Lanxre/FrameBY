package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Address      string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
	
	JWTSecret   string
	StateString string

	DatabaseURL string
	FrontendURL string
	BackendURL  string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return &Config{
		Address:      getEnv("SERVER_ADDRESS", "127.0.0.1:8080"),
		WriteTimeout: getEnvDuration("SERVER_WRITE_TIMEOUT", 15 * time.Second),
		ReadTimeout:  getEnvDuration("SERVER_READ_TIMEOUT", 15 * time.Second),

		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:pass@localhost:5432/mydb?sslmode=disable"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
		BackendURL:  getEnv("BACKEND_URL", "http://localhost:8080"),

		JWTSecret:   getEnv("JWT_SECRET", "jwt_secret"),
		StateString: getEnv("STATE", "random-string-secret"),
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvDuration(key string, defaultVal time.Duration) time.Duration {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultVal
	}

	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Printf("Invalid duration for %s: %s. Using default %v", key, value, defaultVal)
		return defaultVal
	}
	return duration
}
