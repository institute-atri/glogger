package main

import "fmt"

func doPrintbs(a ...interface{}) (str string) {
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
		}
	}

	return
}
