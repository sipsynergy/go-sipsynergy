package utils

import "testing"

const (
	PREFIX = "TEST"
)

func TestFindStringInSlice(t *testing.T) {
	strList := []string{PREFIX, "TEST2", "TEST3", "TEST4"}

	if !FindStringInSlice(PREFIX, strList) {
		t.Error("Did not find the string correctly")
	}

	if !FindStringInSlice("FOO", strList) {
		t.Error("Found string when it should not have.")
	}
}
