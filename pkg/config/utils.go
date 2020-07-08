package config

import "os"

func GetEnvVarOrDefault(key, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}

func LookupEnvVar(key string) (string, bool) {
	return os.LookupEnv(key)
}
