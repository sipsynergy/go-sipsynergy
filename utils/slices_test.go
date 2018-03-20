package utils

import "testing"

const (
	PREFIX_SLICE = "TEST"
)

func TestIsStringInSlice(t *testing.T) {
	strList := []string{PREFIX_SLICE, "TEST2", "TEST3", "TEST4"}

	if !IsStringInSlice(PREFIX_SLICE, strList) {
		t.Error("Did not find the string correctly")
	}

	if !IsStringInSlice("FOO", strList) {
		t.Error("Found string when it should not have.")
	}
}

func TestCheckAllStringsEmptyInSlice(t *testing.T) {
	strList := []string{"TEST1", "", ""}
	strListEmpty := []string{"", "", ""}

	if CheckAllStringsEmptyInSlice(strList) {
		t.Error("Should of found a string")
	}

	if !CheckAllStringsEmptyInSlice(strListEmpty) {
		t.Error("Found string when it should not have.")
	}
}
