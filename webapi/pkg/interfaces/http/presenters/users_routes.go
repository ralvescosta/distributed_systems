package presenters

import (
	"webapi/pkg/app/interfaces"
	adapter "webapi/pkg/infra/adapters"
	server "webapi/pkg/infra/http_server"
	"webapi/pkg/interfaces/http/handlers"
)

type IUsersRoutes interface {
	Register(httpServer server.IHttpServer)
}

type usersRoutes struct {
	handlers handlers.IUsersHandler
	logger   interfaces.ILogger
}

func (pst usersRoutes) Register(httpServer server.IHttpServer) {
	httpServer.RegistreRoute("POST", "/api/v1/users", adapter.HandlerAdapt(pst.handlers.Create, pst.logger))
}

func NewUsersRoutes(logger interfaces.ILogger, handlers handlers.IUsersHandler) IUsersRoutes {
	return usersRoutes{
		handlers,
		logger,
	}
}
