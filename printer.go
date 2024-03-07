package main

import (
	"bufio"
	"fmt"
	"strings"
	"syscall"
)

func Println(a ...interface{}) {
	stdout.WriteString(doPrintbs(a...) + "\n")
}

func Printf(format string, a ...interface{}) {
	stdout.WriteString(fmt.Sprintf(format, a...))
}

func Print(a ...interface{}) {
	stdout.WriteString(doPrintbs(a...))
}

func Fatal(a ...interface{}) {
	stderr.WriteString(doPrintbs(a...))
	syscall.Exit(1)
}

func Done(t ...interface{}) {
	stdout.WriteString(COLORDONE + " " + doPrintbs(t...) + "\n")
}

func Danger(t ...interface{}) {
	stdout.WriteString(COLORDANGER + " " + doPrintbs(t...) + "\n")
}

func Warning(t ...interface{}) {
	stdout.WriteString(COLORWARNING + " " + doPrintbs(t...) + "\n")
}

func Info(t ...interface{}) {
	stdout.WriteString(COLORINFO + " " + doPrintbs(t...) + "\n")
}

func Verbose(t ...interface{}) {
	stdout.WriteString(COLORVERBOSE + " " + doPrintbs(t...) + "\n")
}

func ScanQ(t ...interface{}) (string, error) {
	stdout.WriteString(COLORQUESTION + " " + doPrintbs(t...))

	var response, err = bufio.NewReader(&stdin).ReadString('\n')

	if err != nil {
		return "", err
	}

	response = strings.ToLower(response)

	if response == "\n" {
		return response, nil
	}

	response = strings.Replace(response, "\n", "", -1)

	return response, nil
}
