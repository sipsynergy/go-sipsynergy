package utils

// Supply a string and []string list and it will return a bool if found.
func IsStringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}