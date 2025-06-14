package utils

import (
	"fmt"
	"os"
)

// MakeEnvFunc returns a function that retrieves the value of the specified environment variable.
// If the environment variable is not set, the returned function will panic with an error.
func MakeEnvFunc(envKey string) func() string {
	return func() string {
		var s string
		if val, ok := os.LookupEnv(envKey); ok {
			s = val
		} else {
			panic(fmt.Errorf("%s not set in environment", envKey))
		}

		return s
	}
}

// MakeEnvFuncWithDefault returns a function that retrieves the value of the specified environment variable.
// If the environment variable is not set, the returned function will return the provided default value.
func MakeEnvFuncWithDefault(envKey, defaultVal string) func() string {
	return func() string {
		var s = defaultVal
		if val, ok := os.LookupEnv(envKey); ok {
			s = val
		}

		return s
	}
}
