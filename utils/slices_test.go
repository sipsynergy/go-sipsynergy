package utils

import "testing"

const (
	PREFIX = "TEST"
)

func TestIsStringInSlice(t *testing.T) {
	strList := []string{PREFIX, "TEST2", "TEST3", "TEST4"}

	if !IsStringInSlice(PREFIX, strList) {
		t.Error("Did not find the string correctly")
	}

	if !IsStringInSlice("FOO", strList) {
		t.Error("Found string when it should not have.")
	}
}
