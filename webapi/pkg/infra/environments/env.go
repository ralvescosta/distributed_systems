package environments

import (
	"fmt"
	"os"

	"github.com/ralvescosta/dotenv"
)

var dotEnvConfig = dotenv.Configure

func Configure() error {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "" {
		goEnv = "development"
	}
	return dotEnvConfig(fmt.Sprintf(".env.%s", goEnv))
}
