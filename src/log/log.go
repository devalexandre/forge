package log

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
)

var (
	LogLevel    string // debug, info, warn, error, all
	PrettyPrint bool
)

type LoggerMessage struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

func Infof(format string, args ...interface{}) {
	if LogLevel == "info" || LogLevel == "debug" || LogLevel == "all" {
		msg := prepareLogMessage("info", format, args...)
		color.Blue(msg)
	}
}
func Info(args ...interface{}) {
	if LogLevel == "info" || LogLevel == "debug" || LogLevel == "all" {
		msg := LoggerMessage{
			Level: "info",
			Msg:   fmt.Sprint(args...),
		}
		color.Blue(prettyPrint(msg))

	}
}

func Debugf(format string, args ...interface{}) {
	if LogLevel == "debug" {
		msg := prepareLogMessage("debug", format, args...)
		color.Blue(msg)
	}
}

func Errorf(format string, args ...interface{}) {
	if LogLevel == "error" || LogLevel == "debug" || LogLevel == "all" {
		msg := prepareLogMessage("error", format, args...)
		color.Red(msg)
	}
}

func Warnf(format string, args ...interface{}) {
	if LogLevel == "warn" || LogLevel == "debug" || LogLevel == "all" {
		msg := prepareLogMessage("warn", format, args...)
		color.Yellow(msg)
	}
}

func Successf(format string, args ...interface{}) {
	color.Green(format, args...)
}

func Printf(format string, args ...interface{}) {
	color.White(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	msg := prepareLogMessage("error", format, args...)
	color.Red(msg)
	os.Exit(1)
}

func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func toJSON(i interface{}) string {
	s, _ := json.Marshal(i)
	return string(s)
}

func parseToLogMessage(level string, msg string, args ...interface{}) LoggerMessage {
	return LoggerMessage{
		Level: level,
		Msg:   fmt.Sprintf(msg, args...),
	}
}

func prepareLogMessage(level string, msg string, args ...interface{}) string {
	if !PrettyPrint {
		return toJSON(parseToLogMessage(level, msg, args...))
	}
	return prettyPrint(parseToLogMessage(level, msg, args...))
}
