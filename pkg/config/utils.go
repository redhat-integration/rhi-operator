package config

import "os"

func GetEnvVar(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}
