package cmd

import (
	"os"
	"webapi/pkg/app/interfaces"
	appUseCases "webapi/pkg/app/usecases"
	"webapi/pkg/infra/database"
	"webapi/pkg/infra/hasher"
	httpServer "webapi/pkg/infra/http_server"
	"webapi/pkg/infra/logger"
	"webapi/pkg/infra/repositories"
	tokenManager "webapi/pkg/infra/token_manager"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/middlewares"
	"webapi/pkg/interfaces/http/presenters"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	usersRoutes          presenters.IUsersRoutes
	authenticationRoutes presenters.IAuthenticationRoutes
	inventoryRoutes      presenters.IInventoryRoutes

	monitoring *newrelic.Application
}

func NewContainer() webApiContainer {
	monitoring, err := newrelic.NewApplication(
		newrelic.ConfigAppName(os.Getenv("APP_NAME")),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		// newrelic.ConfigDebugLogger(os.Stdout),
		newrelic.ConfigDistributedTracerEnabled(true),
		func(cfg *newrelic.Config) {
			cfg.CustomInsightsEvents.Enabled = false
			cfg.TransactionTracer.Segments.Threshold = 1
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

	userRepository := repositories.NewUserRepository(logger, dbConnection, monitoring)
	hasher := hasher.NewHahser(logger)
	accessTokenManager := tokenManager.NewTokenManager(logger)
	createUserUseCase := appUseCases.NewCreateUserUseCase(userRepository, hasher, accessTokenManager)
	usersHandler := handlers.NewUsersHandler(logger, createUserUseCase)
	usersRoutes := presenters.NewUsersRoutes(logger, usersHandler)

	authenticationUserUseCase := appUseCases.NewAuthenticateUserUseCase(userRepository, hasher, accessTokenManager)
	authenticationHandler := handlers.NewAuthenticationHandler(logger, authenticationUserUseCase)
	authenticationRoutes := presenters.NewAuthenticationRoutes(logger, authenticationHandler)

	authenticationUseCase := appUseCases.NewAuthenticationUseCase(userRepository, accessTokenManager)
	authenticationMiddleware := middlewares.NewAuthMiddleware(authenticationUseCase)

	inventoryHandler := handlers.NewInventoryHandler(logger)
	inventoryRoutes := presenters.NewInventoryRoutes(logger, authenticationMiddleware, inventoryHandler)

	return webApiContainer{
		logger:     logger,
		httpServer: httpServer,

		usersRoutes:          usersRoutes,
		authenticationRoutes: authenticationRoutes,
		inventoryRoutes:      inventoryRoutes,

		monitoring: monitoring,
	}
}
