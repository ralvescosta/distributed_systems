package cmd

import (
	"os"
	"webapi/pkg/infra/environments"

	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func WebApi() error {
	if err := environments.Configure(); err != nil {
		return err
	}

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		// newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.CustomInsightsEvents.Enabled = true
		},
	)
	if err != nil {
		return err
	}
	container := NewContainer()

	// Server setup
	container.httpServer.Setup()

	//
	container.httpServer.Use(nrgin.Middleware(app))

	// Router register
	container.usersRoutes.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}
	return nil
}
