package logger

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

var level LogLevel

func init() {
	envLogLevel := os.Getenv("LOG_LEVEL")

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
}

func Debugf(format string, args ...interface{}) {
	if level <= LogLevelDebug {
		fmt.Printf(format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if level <= LogLevelInfo {
		fmt.Printf(format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if level <= LogLevelError {
		fmt.Printf(format, args...)
	}
}
