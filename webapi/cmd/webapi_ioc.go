package cmd

import (
	"webapi/pkg/app/interfaces"
	appUseCases "webapi/pkg/app/usecases"
	domainUseCases "webapi/pkg/domain/usecases"
	httpServer "webapi/pkg/infra/http_server"
	"webapi/pkg/infra/logger"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/presenters"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	createUserUseCase domainUseCases.ICreateUserUseCase
	usersHandler      handlers.IUsersHandler
	usersRoutes       presenters.IUsersRoutes
}

func NewContainer() webApiContainer {
	logger := logger.NewLogger()
	httpServer := httpServer.NewHttpServer(logger)

	createUserUseCase := appUseCases.NewCreateUserUseCase()
	usersHandler := handlers.NewUsersHandler(logger, createUserUseCase)
	usersRoutes := presenters.NewUsersRoutes(logger, usersHandler)

	return webApiContainer{
		logger:     logger,
		httpServer: httpServer,

		createUserUseCase: createUserUseCase,
		usersHandler:      usersHandler,
		usersRoutes:       usersRoutes,
	}
}
