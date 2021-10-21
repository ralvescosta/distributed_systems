package handlers

import (
	"webapi/pkg/app/interfaces"
	"webapi/pkg/interfaces/http"
)

type IInventoryHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
}

type inventoryHandler struct {
	logger interfaces.ILogger
}

func (pst inventoryHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	pst.logger.Debug("[InventoryHandler::Create]")
	return http.Created(nil, nil)
}

func NewInventoryHandler(logger interfaces.ILogger) IInventoryHandler {
	return inventoryHandler{
		logger: logger,
	}
}
