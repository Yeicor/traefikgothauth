package traefikoidc

import (
	"fmt"
	"os"
	"time"
)

// Traefik is limited to os.Stdout logging and does noy have the log/slog package for now.

const appName = "traefik-oidc"

func keyValueToString(prefixIfNotEmpty string, keyValue []interface{}) string {
	if len(keyValue) == 0 {
		return ""
	}
	if len(keyValue)%2 != 0 {
		prefixIfNotEmpty += "invalid key-value pairs!"
	}
	for i := 0; i < len(keyValue); i += 2 {
		if i < len(keyValue)-1 {
			prefixIfNotEmpty += fmt.Sprint(keyValue[i]) + "=" + fmt.Sprint(keyValue[i+1]) + " "
		} else {
			prefixIfNotEmpty += fmt.Sprint(keyValue[i])
		}
	}
	return prefixIfNotEmpty
}

// logd logs a message with debug level and key-value pairs.
func logd(msg string, keyValue ...interface{}) {
	_, _ = os.Stdout.WriteString("[" + appName + "] " + time.Now().Format(time.RFC3339) + " DEBUG: " + msg + keyValueToString(" ", keyValue) + "\n")
}

// logi logs a message with info level and key-value pairs.
func logi(msg string, keyValue ...interface{}) {
	_, _ = os.Stdout.WriteString("[" + appName + "] " + time.Now().Format(time.RFC3339) + " INFO: " + msg + keyValueToString(" ", keyValue) + "\n")
}

// loge logs a message with error level and key-value pairs.
func loge(msg string, keyValue ...interface{}) {
	_, _ = os.Stderr.WriteString("[" + appName + "] " + time.Now().Format(time.RFC3339) + " ERROR: " + msg + keyValueToString(" ", keyValue) + "\n")
}
