package cmd

import (
	"fmt"
	"os"
	"webapi/pkg/app/interfaces"
	appUseCases "webapi/pkg/app/usecases"
	"webapi/pkg/infra/database"
	grpcClients "webapi/pkg/infra/grpc_clients"
	"webapi/pkg/infra/hasher"
	httpServer "webapi/pkg/infra/http_server"
	"webapi/pkg/infra/logger"
	"webapi/pkg/infra/repositories"
	"webapi/pkg/infra/telemetry"
	tokenManager "webapi/pkg/infra/token_manager"
	"webapi/pkg/infra/validator"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/middlewares"
	"webapi/pkg/interfaces/http/presenters"
)

type webApiContainer struct {
	logger     interfaces.ILogger
	httpServer httpServer.IHttpServer

	usersRoutes          presenters.IUsersRoutes
	authenticationRoutes presenters.ISessionRoutes
	inventoryRoutes      presenters.IInventoryRoutes
	purchaseRoutes       presenters.IPurchaseRoutes

	telemetryApp telemetry.ITelemetry
}

func NewContainer() webApiContainer {
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
	telemetryApp := telemetry.NewTelemetry()

	userRepository := repositories.NewUserRepository(logger, dbConnection, telemetryApp)
	hasher := hasher.NewHahser(logger)
	accessTokenManager := tokenManager.NewTokenManager(logger)
	createUserUseCase := appUseCases.NewCreateUserUseCase(userRepository, hasher, accessTokenManager)
	usersHandler := handlers.NewUsersHandler(logger, createUserUseCase, validatoR)
	usersRoutes := presenters.NewUsersRoutes(logger, usersHandler)

	authenticationUserUseCase := appUseCases.NewSessionUseCase(userRepository, hasher, accessTokenManager, logger)
	authenticationHandler := handlers.NewSessionHandler(logger, authenticationUserUseCase, validatoR)
	authenticationRoutes := presenters.NewSessionRoutes(logger, authenticationHandler)

	validationTokenUseCase := appUseCases.NewValidatinTokenUseCase(userRepository, accessTokenManager)
	authenticationMiddleware := middlewares.NewAuthMiddleware(validationTokenUseCase)

	inventoryClient := grpcClients.NewInventoryClient(logger, telemetryApp)
	getProductByIdUseCase := appUseCases.NewGetProductByIdUseCase(inventoryClient)
	createProductUseCase := appUseCases.NewCreateProductUseCase(inventoryClient)
	inventoryHandler := handlers.NewInventoryHandler(logger, validatoR, getProductByIdUseCase, createProductUseCase)
	inventoryRoutes := presenters.NewInventoryRoutes(logger, authenticationMiddleware, inventoryHandler)

	pruchaseUseCase := appUseCases.NewPruchaseUseCase()
	purchaseHandler := handlers.NewPurchaseHandler(logger, validatoR, pruchaseUseCase)
	pruchaseRoutes := presenters.NewPruchaseRoutes(purchaseHandler, authenticationMiddleware, logger)

	return webApiContainer{
		logger,
		httpServer,

		usersRoutes,
		authenticationRoutes,
		inventoryRoutes,
		pruchaseRoutes,

		telemetryApp,
	}
}
