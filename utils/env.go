package utils

import (
	"os"
	"strconv"
)

// FromEnv checks the environment for `key`, returning `fallback`
// if not found.
func FromEnv(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

// FromEnvInt checks the enviroment for `key` and converts it to an
// int if found, returning fallback otherwise.
func FromEnvInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	n, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return n
}
