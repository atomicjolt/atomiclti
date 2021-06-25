package config

import (
	"flag"
	"os"
)

func DetermineEnv() string {
	if flag.Lookup("test.v") != nil {
		return "test"
	}

	env := os.Getenv("LTI_ENV")

	if env == "" {
		return "development"
	}

	return env
}
