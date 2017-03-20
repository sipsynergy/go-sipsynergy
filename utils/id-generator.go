package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GenerateHumanID will create a human id.
// It returns the human id string and the associated integer.
func GenerateHumanID(prefix string) string {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(9999999999)
	s := prefix + "-" + strconv.Itoa(i)

	return strings.ToUpper(s)
}
