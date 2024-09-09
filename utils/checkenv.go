package utils

import (
	"fmt"
	"os"
)

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
