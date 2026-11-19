package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func GetEnvString(key, fallback string) string {
	if value := os.Getenv(key); value == "" {
		return fallback
	} else {
		return value
	}
}

func GetDuration(key string, fallback time.Duration) (time.Duration, error) {
	value := os.Getenv(key)

	if value == "" {
		return fallback, nil
	}

	duration, err := time.ParseDuration(value)

	if err != nil {
		log.Fatalf("Invalid %s: %v\n", key, err)
	}

	if duration <= 0 {
		return 0, fmt.Errorf("%s must be positive", key)
	}

	return duration, nil
}
