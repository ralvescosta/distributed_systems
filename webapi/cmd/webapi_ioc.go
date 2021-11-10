package cmd

import (
	"fmt"
	"os"
	"webapi/pkg/app/interfaces"
	appUseCases "webapi/pkg/app/usecases"
	"webapi/pkg/infra/database"
	"webapi/pkg/infra/hasher"
	httpServer "webapi/pkg/infra/http_server"
	"webapi/pkg/infra/logger"
	"webapi/pkg/infra/repositories"
	tokenManager "webapi/pkg/infra/token_manager"
	"webapi/pkg/infra/validator"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/middlewares"
	"webapi/pkg/interfaces/http/presenters"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	usersRoutes          presenters.IUsersRoutes
	authenticationRoutes presenters.ISessionRoutes
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

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	dbConnection, err := database.GetConnection(
		os.Getenv("DB_DRIVER"),
		connectionString,
	)
	if err != nil {
		panic(err)
	}

	logger := logger.NewLogger()
	validatoR := validator.NewValidator()
	httpServer := httpServer.NewHttpServer(logger)

	userRepository := repositories.NewUserRepository(logger, dbConnection)
	hasher := hasher.NewHahser(logger)
	accessTokenManager := tokenManager.NewTokenManager(logger)
	createUserUseCase := appUseCases.NewCreateUserUseCase(userRepository, hasher, accessTokenManager)
	usersHandler := handlers.NewUsersHandler(logger, createUserUseCase, validatoR)
	usersRoutes := presenters.NewUsersRoutes(logger, usersHandler)

	authenticationUserUseCase := appUseCases.NewSessionUseCase(userRepository, hasher, accessTokenManager)
	authenticationHandler := handlers.NewSessionHandler(logger, authenticationUserUseCase, validatoR)
	authenticationRoutes := presenters.NewSessionRoutes(logger, authenticationHandler)

	validationTokenUseCase := appUseCases.NewValidatinTokenUseCase(userRepository, accessTokenManager)
	authenticationMiddleware := middlewares.NewAuthMiddleware(validationTokenUseCase)

	getBookById := appUseCases.NewGetBookByIdUseCase()
	inventoryHandler := handlers.NewInventoryHandler(logger, validatoR, getBookById)
	inventoryRoutes := presenters.NewInventoryRoutes(logger, authenticationMiddleware, inventoryHandler)

	return webApiContainer{
		logger,
		httpServer,

		usersRoutes,
		authenticationRoutes,
		inventoryRoutes,

		monitoring,
	}
}
