package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             string
	DSN              string
	JWTSecret        string
	AllowOrigins     string
	B2KeyID          string
	B2ApplicationKey string
	B2BucketID       string
	B2BucketName     string
}

func Load() *Config {
	_ = godotenv.Load()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Africa/Kinshasa",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "optimatdb"),
		getEnv("DB_SSLMODE", "disable"),
	)

	return &Config{
		Port:             getEnv("PORT", "8080"),
		DSN:              dsn,
		JWTSecret:        getEnv("JWT_SECRET", "change-me-in-production"),
		AllowOrigins:     getEnv("ALLOW_ORIGINS", "http://localhost:4200"),
		B2KeyID:          getEnv("B2_KEY_ID", ""),
		B2ApplicationKey: getEnv("B2_APPLICATION_KEY", ""),
		B2BucketID:       getEnv("B2_BUCKET_ID", ""),
		B2BucketName:     getEnv("B2_BUCKET_NAME", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
