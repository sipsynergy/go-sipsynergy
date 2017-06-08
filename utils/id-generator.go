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

// Generate random string.
func GenerateRandomNumericString(length int) string {
    i := rand.Int63()
    s := strconv.FormatInt(i, 10)[:length]

    return strings.ToUpper(s)
}

func GenerateRandomStirng(length int) string {
    n := length
    b := make([]byte, n)
    if _, err := rand.Read(b); err != nil {
        panic(err)
    }
    s := fmt.Sprintf("%X", b)
    return s
}
