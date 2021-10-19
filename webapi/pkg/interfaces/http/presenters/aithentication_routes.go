package presenters

import (
	"webapi/pkg/app/interfaces"
	adapter "webapi/pkg/infra/adapters"
	server "webapi/pkg/infra/http_server"
	"webapi/pkg/interfaces/http/handlers"
)

type IAuthenticationRoutes interface {
	Register(httpServer server.IHttpServer)
}

type authenticationRoutes struct {
	handlers handlers.IUsersHandler
	logger   interfaces.ILogger
}

func (pst authenticationRoutes) Register(httpServer server.IHttpServer) {
	httpServer.RegistreRoute("POST", "/api/v1/auth", adapter.HandlerAdapt(pst.handlers.Create, pst.logger))
}

func NewAuthenticationRoutes(logger interfaces.ILogger, handlers handlers.IAuthenticationHandler) IAuthenticationRoutes {
	return authenticationRoutes{
		logger:   logger,
		handlers: handlers,
	}
}
