package glogger

// ErrorHandling checks if the given error is not nil.
// If the error is not nil, it logs the error using the Fatal function.
// The Fatal function logs the error to the standard error output (stderr) and exits the program with a status code of 1.
func ErrorHandling(err error) {
	if err != nil {
		Fatal(err)
	}
}

// ErrorWithMessageHandling logs the given error to the standard error output (stderr) along with a custom message.
// If the error is not nil, it uses the Fatal function to log the error.
func ErrorWithMessageHandling(message string, err error) {
	if err != nil {
		Fatal(message, err)
	}
}
