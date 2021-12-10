package presenters

import (
	"webapi/pkg/app/interfaces"
	adapter "webapi/pkg/infra/adapters"
	server "webapi/pkg/infra/http_server"
	"webapi/pkg/interfaces/http/handlers"
	"webapi/pkg/interfaces/http/middlewares"
)

type IPurchaseRoutes interface {
	Register(httpServer server.IHttpServer)
}

type purchaseRoutes struct {
	handlers    handlers.IPurchaseHandler
	middlewares middlewares.IAuthMiddleware
	logger      interfaces.ILogger
}

func (pst purchaseRoutes) Register(httpServer server.IHttpServer) {
	httpServer.RegistreRoute(
		"POST",
		"/api/v1/purchase",
		adapter.MiddlewareAdapt(pst.middlewares.Perform, pst.logger),
		adapter.HandlerAdapt(pst.handlers.Create, pst.logger),
	)
}

func NewPruchaseRoutes(
	handlers handlers.IPurchaseHandler,
	middlewares middlewares.IAuthMiddleware,
	logger interfaces.ILogger,
) IPurchaseRoutes {
	return purchaseRoutes{handlers, middlewares, logger}
}
