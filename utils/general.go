package utils

// PanicOnError logs error message and terminates the main process.
func PanicOnError(err error, message string) {
	if err != nil {
		HandleError(err, message)
	}
}
