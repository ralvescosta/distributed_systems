package cmd

import (
	"os"
	"webapi/pkg/app/interfaces"
	appUseCases "webapi/pkg/app/usecases"
	domainUseCases "webapi/pkg/domain/usecases"
	"webapi/pkg/infra/database"
	httpServer "webapi/pkg/infra/http_server"
	"webapi/pkg/infra/logger"
	"webapi/pkg/infra/repositories"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/presenters"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	createUserUseCase domainUseCases.ICreateUserUseCase
	usersHandler      handlers.IUsersHandler
	usersRoutes       presenters.IUsersRoutes

	monitoring *newrelic.Application
}

func NewContainer() webApiContainer {
	monitoring, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		// newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.CustomInsightsEvents.Enabled = true
		},
	)
	if err != nil {
		panic(err)
	}

	dbConnection, err := database.GetConnection(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger()
	httpServer := httpServer.NewHttpServer(logger)

	userRepository := repositories.NewUserRepository(dbConnection, monitoring)
	createUserUseCase := appUseCases.NewCreateUserUseCase(userRepository)
	usersHandler := handlers.NewUsersHandler(logger, createUserUseCase)
	usersRoutes := presenters.NewUsersRoutes(logger, usersHandler)

	return webApiContainer{
		logger:     logger,
		httpServer: httpServer,

		createUserUseCase: createUserUseCase,
		usersHandler:      usersHandler,
		usersRoutes:       usersRoutes,

		monitoring: monitoring,
	}
}
