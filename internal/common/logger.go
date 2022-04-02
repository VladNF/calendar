package common

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

func NewLogger(level, filename string) Logger {
	log := logrus.New()
	if l, err := logrus.ParseLevel(strings.ToLower(level)); err == nil {
		log.SetLevel(l)
	} else {
		log.Infof("invalid log level passed '%v', using default INFO instead\n", level)
		log.SetLevel(logrus.InfoLevel)
	}

	if f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0o755); err != nil {
		log.Warnf("invalid log file path '%v', logging to Stderr instead", filename)
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(f)
	}
	return log
}
