package utils

import "os"

// GetEnv will get the environment value from OS, if empty will use the fallback given.
// returns string.
func GetEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
