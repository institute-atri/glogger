package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
)

const (
	DONE    = GREEN + "[DONE]" + RESET
	DANGER  = RED + "[DANGER]" + RESET
	FATAL   = RED + "[ERROR]" + RESET
	INFO    = MAGENTA + "[INFO]" + RESET
	WARNING = YELLOW + "[WARN]" + RESET
	SCAN    = YELLOW + "[QUEST]" + RESET
	VERBOSE = BLUE + "[VERBOSE]" + RESET
)

var (
	stdin  = *os.Stdin
	stdout = *os.Stdout
	stderr = *os.Stderr
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
	stdout.WriteString(DONE + " " + doPrintbs(t...) + "\n")
}

func Danger(t ...interface{}) {
	stdout.WriteString(DANGER + " " + doPrintbs(t...) + "\n")
}

func Warning(t ...interface{}) {
	stdout.WriteString(WARNING + " " + doPrintbs(t...) + "\n")
}

func Info(t ...interface{}) {
	stdout.WriteString(INFO + " " + doPrintbs(t...) + "\n")
}

func Verbose(t ...interface{}) {
	stdout.WriteString(VERBOSE + " " + doPrintbs(t...) + "\n")
}

func ScanQ(t ...interface{}) (string, error) {
	stdout.WriteString(SCAN + " " + doPrintbs(t...))

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
