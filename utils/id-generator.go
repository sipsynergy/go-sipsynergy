package utils

import (
	"strings"
	"math/rand"
	"strconv"
)

func GenerateId(prefix string) string {

	i := rand.Intn(9999999999)

	s := prefix + "-" + strconv.Itoa(i)

	return strings.ToUpper(s)
}