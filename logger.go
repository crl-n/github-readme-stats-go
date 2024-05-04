package main

import (
	"fmt"
	"os"
)

type LogLevel int

const (
	LogLevelDebug LogLevel = iota
	LogLevelInfo
	LogLevelError
)

type Logger struct {
	LogLevel
}

func newLogger() *Logger {
	envLogLevel := os.Getenv("LOG_LEVEL")

	var level LogLevel
	switch envLogLevel {
	case "DEBUG":
		level = LogLevelDebug
	case "INFO":
		level = LogLevelInfo
	case "ERROR":
		level = LogLevelError
	default:
		level = LogLevelInfo
	}

	return &Logger{
		LogLevel: level,
	}
}

// Create global logger
var logger = newLogger()

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.LogLevel <= LogLevelDebug {
		fmt.Printf(format, args...)
	}
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.LogLevel <= LogLevelError {
		fmt.Printf(format, args...)
	}
}
