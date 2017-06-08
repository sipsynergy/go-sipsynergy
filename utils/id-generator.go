package utils

import (
    "fmt"
    "math/rand"
    "regexp"
    "strconv"
    "strings"
)

// GenerateHumanID will create a human id.
// It returns the human id string.
func GenerateHumanID(prefix string) string {
    r := GenerateRandomNumericString(10)
    s := prefix + "-" + r

    return strings.ToUpper(s)
}

// ValidateHumanID validates the given string.
// returns true or false.
func ValidateHumanID(humanID string) bool {
    re := regexp.MustCompile(`^[A-Z]+-\d+`)

    return re.MatchString(humanID)
}

// Generate random string.
func GenerateRandomNumericString(length int) string {
    i := rand.Int63()
    s := strconv.FormatInt(i, 10)[:length]

    return strings.ToUpper(s)
}

// Generate random string.
func GenerateRandomString(length int) (string, error) {
    n := length
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil {
        return nil, err
    }
    s := fmt.Sprintf("%X", b)
    return s, nil
}
