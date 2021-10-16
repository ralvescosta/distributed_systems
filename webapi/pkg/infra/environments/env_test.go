package environments

import (
	"os"
	"testing"
)

func Test_Configure_Env_Without_GoEnv(t *testing.T) {
	os.Setenv("GO_ENV", "")
	var envFile string
	dotEnvConfig = func(arg string) error {
		envFile = arg
		return nil
	}

	err := Configure()

	if err != nil {
		t.Error("dotenv need to call only dotEnvConfig function")
	}
	if envFile != ".env.development" {
		t.Error("dotenv need to call only dotEnvConfig function")
	}
}

func Test_Configure_Env_With_GoEnv(t *testing.T) {
	os.Setenv("GO_ENV", "production")
	var envFile string
	dotEnvConfig = func(arg string) error {
		envFile = arg
		return nil
	}

	err := Configure()

	if err != nil {
		t.Error("dotenv need to call only dotEnvConfig function")
	}
	if envFile != ".env.production" {
		t.Error("dotenv need to call only dotEnvConfig function")
	}
}
