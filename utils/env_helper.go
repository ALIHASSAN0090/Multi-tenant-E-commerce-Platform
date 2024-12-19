package utils

import (
	"os"
	"strings"
)

func GetEnvValue(envTag, defaultTag string) string {
	value := strings.TrimSpace(os.Getenv(envTag))
	if value == "" && defaultTag != "" {
		value = defaultTag
	}
	return value
}
