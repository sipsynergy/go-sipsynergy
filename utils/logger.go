package utils

import "log"

// HandleError will log the given error message to Stdout.
func HandleError(err error, message string) {
	log.Println(message)
	log.Panic(err.Error())
}
