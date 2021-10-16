package main

import (
	"fmt"
	"os"

	newrelic "github.com/newrelic/go-agent/v3/newrelic"
)

func main() {

	_, err := newrelic.NewApplication(
		newrelic.ConfigAppName("DS_WebApi"),
		newrelic.ConfigLicense("__YOUR_NEW_RELIC_LICENSE_KEY__"),
		newrelic.ConfigDebugLogger(os.Stdout),
		func(cfg *newrelic.Config) {
			cfg.CustomInsightsEvents.Enabled = false
		})
	if err != nil {
		panic(err)
	}

	fmt.Println("Hello World!")
}
