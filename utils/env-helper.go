package utils

import "os"

// GetEnv will get the environment value from OS, if empty will use the fallback given.
// returns interface.
func GetEnv(key string, fallback interface{}) interface{} {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
