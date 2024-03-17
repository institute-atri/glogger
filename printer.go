package glogger

import (
	"bufio"
	"fmt"
	"strings"
	"syscall"
)

// Println prints a space-separated list of values to standard output followed by a newline character.
// It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func Println(a ...interface{}) {
	_, err := stdout.WriteString(formatArgs(a...) + "\n")
	ErrorHandling(err)
}

// Printf formats and prints the values according to the format string.
// It writes the formatted string to standard output.
// The format string specifies how the later arguments are formatted and printed.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func Printf(format string, a ...interface{}) {
	_, err := stdout.WriteString(fmt.Sprintf(format, a...))
	ErrorHandling(err)
}

// Print prints a space-separated list of values to standard output.
// It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func Print(a ...interface{}) {
	_, err := stdout.WriteString(formatArgs(a...))
	ErrorHandling(err)
}

// Fatal writes a space-separated list of values to standard error without appending a newline character.
// It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the error output, it calls the ErrorHandling function to handle the error.
// Once the error is handled, it exits the program with a status of -1 using the syscall.Exit function.
func Fatal(a ...interface{}) {
	_, err := stderr.WriteString(formatArgs(a...))
	ErrorHandling(err)
	syscall.Exit(-1)
}

// Done writes a space-separated list of values to standard output, prefixed with the COLORDONE constant,
// followed by a newline character. It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func Done(t ...interface{}) {
	_, err := stdout.WriteString(COLORDONE + " " + formatArgs(t...) + "\n")
	ErrorHandling(err)
}

// Danger writes a space-separated list of values to standard output in danger color followed by a newline character.
// It uses the COLORDANGER constant to add the danger color to the output,
// and the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
//
// stdout is an os.File variable representing the standard output.
// COLORDANGER is a string constant representing the color to be added to the output.
// formatArgs is a function that formats the input values.
// ErrorHandling is a function that handles errors.
// stderr is an os.File variable representing the standard error.
func Danger(t ...interface{}) {
	_, err := stdout.WriteString(COLORDANGER + " " + formatArgs(t...) + "\n")
	ErrorHandling(err)
}

// Warning writes a space-separated list of values to standard output with a "WARNING" Prefix and yellow color, followed by a newline character.
// It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func Warning(t ...interface{}) {
	_, err := stdout.WriteString(COLORWARNING + " " + formatArgs(t...) + "\n")
	ErrorHandling(err)
}

// Info prints a space-separated list of values to standard output preceded by the string "[INFO]" in magenta color, followed by a newline character.
// It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func Info(t ...interface{}) {
	_, err := stdout.WriteString(COLORINFO + " " + formatArgs(t...) + "\n")
	ErrorHandling(err)
}

// Verbose prints a space-separated list of values to standard output followed by a newline character, with a "[VERBOSE]" Prefix in blue color.
// It uses the formatArgs function to format the values.
// The values can be of type bool, string, error, int, float64, or float32.
// If an error occurs while writing to the output, it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
// Verbose is designed to provide verbose information to the user.
func Verbose(t ...interface{}) {
	_, err := stdout.WriteString(COLORVERBOSE + " " + formatArgs(t...) + "\n")
	ErrorHandling(err)
}

// ScanQ reads a line of user input from the standard input and returns it as a string.
// It displays a formatted question prompt using the COLORQUESTION constant.
// The question prompt is created by combining the COLORQUESTION constant and the formatted arguments passed in the 't' parameter.
// It uses the formatArgs function to format the question prompt.
// The user's response is read using bufio.NewReader from the standard input until a newline character is encountered.
// The response is then converted to lowercase and any trailing newline characters are removed.
// If the user enters only a newline character, an empty string is returned.
// Otherwise, the response string is returned.
// If an error occurs while writing to the standard output or reading from the standard input,
// it calls the ErrorHandling function to handle the error.
// If the error is not nil, it calls the Fatal function to write the error to standard error and exit the program with a status of -1.
func ScanQ(t ...interface{}) string {
	_, err := stdout.WriteString(COLORQUESTION + " " + formatArgs(t...))
	ErrorHandling(err)

	response, err := bufio.NewReader(&stdin).ReadString('\n')
	ErrorHandling(err)

	response = strings.ToLower(response)

	if response == "\n" {
		return response
	}

	response = strings.Replace(response, "\n", "", -1)
	return response
}
