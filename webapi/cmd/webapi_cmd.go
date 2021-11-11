package cmd

import (
	"webapi/pkg/infra/environments"
)

func WebApi() error {
	if err := environments.Configure(); err != nil {
		return err
	}

	container := NewContainer()

	// Server setup
	container.httpServer.Setup()

	//middlewares
	container.httpServer.RegisterMiddleware(container.telemetryApp.GinMiddle())

	// Router register
	container.usersRoutes.Register(container.httpServer)
	container.authenticationRoutes.Register(container.httpServer)
	container.inventoryRoutes.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}
	return nil
}
