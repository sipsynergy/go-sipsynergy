package utils

import "testing"

const (
	PREFIX = "TEST"
)

func TestGenerateHumanID(t *testing.T) {
	humanID := GenerateHumanID(PREFIX)

	if !ValidateHumanID(humanID) {
		t.Error("Failed generating human id. " + humanID)
	}
}

func TestValidateHumanID(t *testing.T) {
	validHumanID := GenerateHumanID(PREFIX)
	invalidHumanID := "IAHDKASD"

	if !ValidateHumanID(validHumanID) {
		t.Error("Generated a bad human id.")
	}

	if ValidateHumanID(invalidHumanID) {
		t.Error("Failed to validate the human id correctly.")
	}
}
