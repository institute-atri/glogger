package main

import (
	"os"
	"time"
)

type logger struct {
	prefix  string
	silient bool
}

func NewLogger(prefix string, silient bool) *logger {
	return &logger{prefix: prefix + ":", silient: silient}
}

func (l *logger) Done(s string) error {
	var message = DONE + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.prefix + " " + s + "\n"

	var err = l.WriterLogsFile(message)

	if err != nil {
		return err
	}

	if !l.silient {
		stdout.WriteString(message)
	}

	return nil
}

func (l *logger) Danger(s string) error {
	var message = DANGER + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.prefix + " " + s + "\n"

	var err = l.WriterLogsFile(message)

	if err != nil {
		return err
	}

	if !l.silient {
		stdout.WriteString(message)
	}

	return nil
}

func (l *logger) Warning(s string) error {
	var message = WARNING + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.prefix + " " + s + "\n"

	var err = l.WriterLogsFile(message)

	if err != nil {
		return err
	}

	if !l.silient {
		stdout.WriteString(message)
	}
	return nil
}

func (l *logger) Info(s string) error {
	var message = INFO + " " + time.Now().Format("02/01/2006 15:04:05") + " " + l.prefix + " " + s + "\n"

	var err = l.WriterLogsFile(message)

	if err != nil {
		return err
	}

	if !l.silient {
		stdout.WriteString(message)
	}

	return nil
}

func (l *logger) WriterLogsFile(s string) error {
	var logs = "logs"

	var _, err = os.Stat(logs)

	if os.IsNotExist(err) {
		var archive, err = os.Create(logs)

		if err != nil {
			return err
		}

		archive.Close()
	} else if err != nil {
		return err
	}

	file, err := os.OpenFile(logs, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	file.WriteString(s)

	defer file.Close()

	return nil
}
