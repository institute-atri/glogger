package glogger

import "fmt"

// formatArgs formats the provided variadic arguments into a single string.
// Each argument is converted to a string representation based on its type.
// The resulting string is constructed by appending each argument's string representation,
// separated by a space.
//
// The following types are supported:
//   - bool: converted using fmt.Sprint
//   - string: appended as is
//   - error: converted to string using Error() method
//   - int: converted using fmt.Sprint
//   - float64: converted using fmt.Sprint
//   - float32: converted using fmt.Sprint
//
// Example usage:
//
//	formatArgs(true, "hello", fmt.Errorf("error"), 42, 3.14)
//	// Output: "true hello error 42 3.14"
func formatArgs(a ...interface{}) (str string) {
	for count, arg := range a {
		if count > 0 {
			str += " "
		}
		switch arg.(type) {
		case bool:
			str += fmt.Sprint(arg.(bool))
		case string:
			str += arg.(string)
		case error:
			str += arg.(error).Error()
		case int:
			str += fmt.Sprint(arg.(int))
		case float64:
			str += fmt.Sprint(arg.(float64))
		case float32:
			str += fmt.Sprint(arg.(float32))
		}
	}
	return
}
