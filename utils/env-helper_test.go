package utils

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	str := GetEnv("test", "fallback")

	if len(str) == 0{
		t.Error("No fallback was supplied for missing os env")
	}
}
