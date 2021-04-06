package env

import (
	"os"
	"strconv"
)

func GetEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	} else {
		return defaultVal
	}
}

func GetEnvString(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	} else {
		return defaultVal
	}
}

func GetEnvInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		if d, err := strconv.ParseInt(val, 10, 64); err == nil {
			return int(d)
		}
	}

	return defaultVal
}
