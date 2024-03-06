package main

import (
	"os"
	"time"
)

var (
	LDONE    = "[DONE] " + time.Now().Format("02/01/2006 15:04:05")
	LDANGER  = "[DANGER] " + time.Now().Format("02/01/2006 15:04:05")
	LFATAL   = "[FATAL] " + time.Now().Format("02/01/2006 15:04:05")
	LINFO    = "[INFO] " + time.Now().Format("02/01/2006 15:04:05")
	LWARNING = "[WARN] " + time.Now().Format("02/01/2006 15:04:05")
)

type logger struct {
	prefix  string
	silient bool
}

func NewLogger(prefix string, silient bool) *logger {
	var l = logger{
		prefix:  prefix + " :",
		silient: silient,
	}

	return &l
}

func (l *logger) Done(t ...interface{}) {
	l.WriterLogsFile(LDONE + " " + l.prefix + " " + doPrintbs(t...) + "\n")

	if !l.silient {
		stdout.WriteString(LDONE + " " + l.prefix + " " + doPrintbs(t...) + "\n")
	}
}

func (l *logger) Danger(t ...interface{}) {
	l.WriterLogsFile(LDANGER + " " + l.prefix + " " + doPrintbs(t...) + "\n")

	if !l.silient {
		stdout.WriteString(LDANGER + " " + l.prefix + " " + doPrintbs(t...) + "\n")
	}
}

func (l *logger) Warning(t ...interface{}) {
	l.WriterLogsFile(LWARNING + " " + l.prefix + " " + doPrintbs(t...) + "\n")

	if !l.silient {
		stdout.WriteString(LWARNING + " " + l.prefix + " " + doPrintbs(t...) + "\n")
	}
}

func (l *logger) Info(t ...interface{}) {
	l.WriterLogsFile(LINFO + " " + l.prefix + " " + doPrintbs(t...) + "\n")

	if !l.silient {
		stdout.WriteString(LINFO + " " + l.prefix + " " + doPrintbs(t...) + "\n")
	}
}

func (l *logger) WriterLogsFile(s string) error {
	var filelogs = "logs"

	var _, err = os.Stat(filelogs)

	if os.IsNotExist(err) {
		var archive, err = os.Create(filelogs)

		if err != nil {
			return err
		}

		archive.Close()
	} else if err != nil {
		return err
	}

	file, err := os.OpenFile(filelogs, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return err
	}

	file.WriteString(s)

	defer file.Close()

	return nil
}
