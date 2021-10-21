package cmd

import (
	"webapi/pkg/infra/environments"

	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func WebApi() error {
	if err := environments.Configure(); err != nil {
		return err
	}

	container := NewContainer()

	// Server setup
	container.httpServer.Setup()

	//middlewares
	container.httpServer.RegisterMiddleware(nrgin.Middleware(container.monitoring))

	// Router register
	container.usersRoutes.Register(container.httpServer)
	container.authenticationRoutes.Register(container.httpServer)
	container.inventoryRoutes.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}
	return nil
}
