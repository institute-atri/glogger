package glogger

import (
	"os"
	"time"
)

// Logger is a type that handles logging messages with a specified prefix and specific log levels.
type Logger struct {
	Prefix   string
	IsSilent bool
}

// NewLogger returns a pointer to a new instance of the Logger struct.
// It takes in a prefix string and a silent bool parameter.
// The prefix string is used to add a prefix to the log messages.
// The silent bool parameter determines whether the log messages should also be written to stdout.
// The returned Logger struct has methods for logging messages of different levels (Done, Danger, Warning, and Info).
// The methods use the prefix and silent parameters passed to NewLogger to format the log messages.
// The logs are also written to a file named "logs" in the same directory.
// If the "logs" file does not exist, it is created.
// The Logger struct returned by NewLogger satisfies the io.Writer interface,
// allowing the log messages to be written to stdout using the fmt.Fprint function.
// The log messages are formatted with the current timestamp.
func NewLogger(prefix string, silient bool) *Logger {
	return &Logger{
		Prefix:   prefix + ":",
		IsSilent: silient,
	}
}

// Done outputs a log message with the "[DONE]" prefix, the current timestamp, the logger prefix, and the provided message.
// It writes the log message to a log file and to stdout if IsSilent is false.
func (l *Logger) Done(s string) error {
	message := DONE + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.Prefix + " " + s + "\n"

	err := l.WriterLogsFile(message)
	ErrorWithMessageHandling(message, err)

	if !l.IsSilent {
		_, err := stdout.WriteString(message)
		ErrorWithMessageHandling(message, err)
	}

	return nil
}

// Danger outputs a log message with the "[DANGER]" prefix, the current timestamp, the logger prefix, and the provided message.
// It writes the log message to a log file and to stdout if IsSilent is false.
func (l *Logger) Danger(s string) error {
	message := DANGER + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.Prefix + " " + s + "\n"

	err := l.WriterLogsFile(message)
	ErrorWithMessageHandling(message, err)

	if !l.IsSilent {
		_, err := stdout.WriteString(message)
		ErrorWithMessageHandling(message, err)
	}

	return nil
}

// Warning outputs a log message with the "[WARNING]" prefix, the current timestamp, the logger prefix, and the provided message.
// It writes the log message to a log file and to stdout if IsSilent is false.
func (l *Logger) Warning(s string) error {
	var message = WARNING + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.Prefix + " " + s + "\n"

	err := l.WriterLogsFile(message)
	ErrorWithMessageHandling(message, err)

	if !l.IsSilent {
		_, err := stdout.WriteString(message)
		ErrorWithMessageHandling(message, err)
	}
	return nil
}

// Info outputs a log message with the "[INFO]" prefix, the current timestamp, the logger prefix, and the provided message.
// It writes the log message to a log file and to stdout if IsSilent is false.
func (l *Logger) Info(s string) error {
	message := INFO + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.Prefix + " " + s + "\n"

	err := l.WriterLogsFile(message)
	ErrorWithMessageHandling(message, err)

	if !l.IsSilent {
		_, err := stdout.WriteString(message)
		ErrorWithMessageHandling(message, err)
	}

	return nil
}

// WriterLogsFile writes the provided string to a log file.
// The log file is created if it does not exist.
// If the log file exists, the string is appended to the file.
// Any errors encountered during file operations are handled by the ErrorHandling function.
func (l *Logger) WriterLogsFile(s string) error {
	logs := "logs"

	_, err := os.Stat(logs)
	ErrorHandling(err)

	if os.IsNotExist(err) {
		archive, err := os.Create(logs)
		ErrorHandling(err)

		err = archive.Close()
		ErrorHandling(err)

	}

	file, err := os.OpenFile(logs, os.O_WRONLY|os.O_APPEND, 0644)
	ErrorHandling(err)

	_, err = file.WriteString(s)
	ErrorHandling(err)

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			ErrorHandling(err)
		}
	}(file)

	return nil
}
