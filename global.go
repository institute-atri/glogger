package glogger

import "os"

/*
	In this specific file, only declarations of global scope constants and variables will be accepted,
	while functions, structs, and interfaces will not be deemed appropriate.
*/

const (
	DONE     = "[DONE]"
	DANGER   = "[DANGER]"
	FATAL    = "[ERROR]"
	INFO     = "[INFO]"
	WARNING  = "[WARNING]"
	QUESTION = "[QUEST]"
	VERBOSE  = "[VERBOSE]"
)

const (
	RED       = "\u001b[31;1m"
	BLUE      = "\u001b[34;1m"
	GREEN     = "\u001b[32;1m"
	BLACK     = "\u001b[30;1m"
	WHITE     = "\u001b[37;1m"
	YELLOW    = "\u001b[33;1m"
	MAGENTA   = "\u001b[35;1m"
	CYAN      = "\u001b[36;1m"
	RESET     = "\u001b[0m"
	BOLD      = "\u001b[1m"
	UNDERLINE = "\u001b[4m"
	REVERSED  = "\u001b[7m"
)

const (
	COLORDONE     = GREEN + "[DONE]" + RESET
	COLORDANGER   = RED + "[DANGER]" + RESET
	COLORFATAL    = RED + "[ERROR]" + RESET
	COLORINFO     = MAGENTA + "[INFO]" + RESET
	COLORWARNING  = YELLOW + "[WARNING]" + RESET
	COLORQUESTION = YELLOW + "[QUESTION]" + RESET
	COLORVERBOSE  = BLUE + "[VERBOSE]" + RESET
)

var (
	stdin  = *os.Stdin
	stdout = *os.Stdout
	stderr = *os.Stderr
)
