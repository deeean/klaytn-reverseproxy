package utils

import (
	"net/url"
	"os"
)

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetEnvURLOrDefault(key string, defaultValue string) (*url.URL, error) {
	value := GetEnvOrDefault(key, defaultValue)
	u, err := url.Parse(value)
	if err != nil {
		return nil, err
	}

	return u, nil
}
