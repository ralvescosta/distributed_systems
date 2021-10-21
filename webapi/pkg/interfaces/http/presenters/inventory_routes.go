package presenters

import (
	"webapi/pkg/app/interfaces"
	adapter "webapi/pkg/infra/adapters"
	server "webapi/pkg/infra/http_server"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/middlewares"
)

type IInventoryRoutes interface {
	Register(httpServer server.IHttpServer)
}

type inventoryRoutes struct {
	handlers    handlers.IInventoryHandler
	middlewares middlewares.IAuthMiddleware
	logger      interfaces.ILogger
}

func (pst inventoryRoutes) Register(httpServer server.IHttpServer) {
	httpServer.RegistreRoute(
		"GET",
		"/api/v1/inventory",
		adapter.MiddlewareAdapt(pst.middlewares.Perform, pst.logger),
		adapter.HandlerAdapt(pst.handlers.Create, pst.logger),
	)
}

func NewInventoryRoutes(logger interfaces.ILogger, middlewares middlewares.IAuthMiddleware, handlers handlers.IInventoryHandler) IInventoryRoutes {
	return inventoryRoutes{
		logger:      logger,
		middlewares: middlewares,
		handlers:    handlers,
	}
}
