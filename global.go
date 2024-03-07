package main

import "os"

/* In this specific file, only declarations of global scope constants and variables will be accepted, while functions, structs, and interfaces will not be deemed appropriate. */

const (
	DONE     = "[DONE]"
	DANGER   = "[DANGER]"
	FATAL    = "[ERROR]"
	INFO     = "[INFO]"
	WARNING  = "[WARNG]"
	QUESTION = "[QUEST]"
	VERBOSE  = "[VERBOSE]"
)

var (
	stdin  = *os.Stdin
	stdout = *os.Stdout
	stderr = *os.Stderr
)
