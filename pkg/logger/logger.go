package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

type Logger struct {
	*log.Logger
}

func NewLogger() *Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})
	
	return &Logger{
		Logger: logger,
	}
}

func (l *Logger) Info(msg string, args ...interface{}) {
	l.Logger.Helper()
	l.Logger.Infof(msg, args...)
}

func (l *Logger) Error(err error, msg string, args ...interface{}) {
	l.Logger.Helper()
	if err != nil {
		l.Logger.Error(msg, "error", err, "details", args)
	} else {
		l.Logger.Errorf(msg, args...)
	}
}

func (l *Logger) Fatal(err error, msg string, args ...interface{}) {
	l.Logger.Helper()
	if err != nil {
		l.Logger.Fatal(msg, "error", err, "details", args)
	} else {
		l.Logger.Fatalf(msg, args...)
	}
}

var DefaultLogger = NewLogger()

func Info(msg string, args ...interface{}) {
	DefaultLogger.Helper()
	DefaultLogger.Info(msg, args...)
}

func Error(err error, msg string, args ...interface{}) {
	DefaultLogger.Helper()
	DefaultLogger.Error(err, msg, args...)
}

func Fatal(err error, msg string, args ...interface{}) {
	DefaultLogger.Helper()
	DefaultLogger.Fatal(err, msg, args...)
}
