package presenters

import (
	"webapi/pkg/app/interfaces"
	adapter "webapi/pkg/infra/adapters"
	server "webapi/pkg/infra/http_server"
	"webapi/pkg/interfaces/http/handlers"
)

type ISessionRoutes interface {
	Register(httpServer server.IHttpServer)
}

type sessionRoutes struct {
	handlers handlers.IUsersHandler
	logger   interfaces.ILogger
}

func (pst sessionRoutes) Register(httpServer server.IHttpServer) {
	httpServer.RegistreRoute("POST", "/api/v1/auth", adapter.HandlerAdapt(pst.handlers.Create, pst.logger))
}

func NewSessionRoutes(logger interfaces.ILogger, handlers handlers.ISessionHandler) ISessionRoutes {
	return sessionRoutes{
		handlers,
		logger,
	}
}
