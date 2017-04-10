package utils

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	str := GetEnv("getenvtest", "fallback")

	if len(str) == 0{
		t.Error("No fallback was supplied for missing os env")
	}

	if str != "fallback" {
		t.Error("Different fallback string was given")
	}

	os.Setenv("ValueTest", "test123")

	str = GetEnv("ValueTest", "")

	if len(str) == 0 {
		t.Error("Could not retrieve the env")
	}

	if str != "test123" {
		t.Error("Wrong env retrieved or fallback used")
	}
}
