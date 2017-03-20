package utils

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
)

// GenerateHumanID will create a human id.
// It returns the human id string.
func GenerateHumanID(prefix string) string {
	i := rand.Int63()
	s := prefix + "-" + strconv.FormatInt(i, 10)[:10]

	return strings.ToUpper(s)
}

// ValidateHumanID validates the given string.
// returns true or false.
func ValidateHumanID(humanID string) bool {
	re := regexp.MustCompile(`^[A-Z]+-\d+`)

	return re.MatchString(humanID)
}
