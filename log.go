package traefikgothauth

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Traefik is limited to os.Stdout logging and does noy have the log/slog package for now.

const appName = "traefik-oidc"

// LogLevel is the log level.
type LogLevel int

var logLevelCurrent = logLevelTrace

const (
	logLevelTrace LogLevel = 0
	logLevelDebug LogLevel = 1
	logLevelInfo  LogLevel = 2
	logLevelWarn  LogLevel = 3
	logLevelError LogLevel = 4
	logLevelOff   LogLevel = 5
)

var logLevelText = map[LogLevel]string{
	logLevelTrace: "TRACE",
	logLevelDebug: "DEBUG",
	logLevelInfo:  "INFO ",
	logLevelWarn:  "WARN ",
	logLevelError: "ERROR",
	logLevelOff:   "OFF  ",
}

var logTextLevel = func() map[string]LogLevel {
	m := make(map[string]LogLevel)
	for k, v := range logLevelText {
		m[strings.ReplaceAll(v, " ", "")] = k
	}
	return m
}()

func keyValueToString(prefixIfNotEmpty string, keyValue []interface{}) string {
	if len(keyValue) == 0 {
		return ""
	}
	if len(keyValue)%2 != 0 {
		prefixIfNotEmpty += "invalid key-value pairs! "
	}
	for i := 0; i < len(keyValue); i += 2 {
		if i < len(keyValue)-1 {
			prefixIfNotEmpty += fmt.Sprintf("%+v=%+v ", keyValue[i], keyValue[i+1])
		} else {
			prefixIfNotEmpty += fmt.Sprintf("%+v ", keyValue[i])
		}
	}
	return prefixIfNotEmpty
}

const timeLayout = time.DateTime + " -0700"

func logFormat(level LogLevel, msg string, keyValue []interface{}) string {
	return "[" + appName + "] [" + logLevelText[level] + "] " + time.Now().Format(timeLayout) + " | " + msg + keyValueToString(" ", keyValue) + "\n"
}

// logt logs a message with trace level and key-value pairs.
func logt(msg string, keyValue ...interface{}) {
	if !logtEnabled() {
		return
	}
	_, _ = os.Stdout.Write([]byte(logFormat(logLevelTrace, msg, keyValue)))
}

func logtEnabled() bool {
	return logLevelCurrent <= logLevelTrace
}

// logd logs a message with debug level and key-value pairs.
func logd(msg string, keyValue ...interface{}) {
	if !logdEnabled() {
		return
	}
	_, _ = os.Stdout.Write([]byte(logFormat(logLevelDebug, msg, keyValue)))
}

func logdEnabled() bool {
	return logLevelCurrent <= logLevelDebug
}

// logi logs a message with info level and key-value pairs.
func logi(msg string, keyValue ...interface{}) {
	if !logiEnabled() {
		return
	}
	_, _ = os.Stdout.Write([]byte(logFormat(logLevelInfo, msg, keyValue)))
}

func logiEnabled() bool {
	return logLevelCurrent <= logLevelInfo
}

// logw logs a message with warn level and key-value pairs.
func logw(msg string, keyValue ...interface{}) {
	if !logwEnabled() {
		return
	}
	_, _ = os.Stderr.Write([]byte(logFormat(logLevelWarn, msg, keyValue)))
}

func logwEnabled() bool {
	return logLevelCurrent <= logLevelWarn
}

// loge logs a message with error level and key-value pairs.
func loge(msg string, keyValue ...interface{}) {
	if !logeEnabled() {
		return
	}
	_, _ = os.Stderr.Write([]byte(logFormat(logLevelError, msg, keyValue)))
}

func logeEnabled() bool {
	return logLevelCurrent <= logLevelError
}
