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

// CheckAllStringsEmptyInSlice takes a []string and it will return true if all strings are empty.
func CheckAllStringsEmptyInSlice(a []string) bool {
    for i := 1; i < len(a); i++ {
        if a[i] != "" {
            return false
        }
    }
    return true
}