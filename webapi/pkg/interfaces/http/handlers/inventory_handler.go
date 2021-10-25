package handlers

import (
	"webapi/pkg/app/interfaces"
	"webapi/pkg/interfaces/http"
)

type IInventoryHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
	GetById(httpRequest http.HttpRequest) http.HttpResponse
}

type inventoryHandler struct {
	logger    interfaces.ILogger
	validator interfaces.IValidator
}

func (pst inventoryHandler) GetById(httpRequest http.HttpRequest) http.HttpResponse {
	pst.logger.Debug("[InventoryHandler::GetById]")
	return http.Created(nil, nil)
}

func (pst inventoryHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	pst.logger.Debug("[InventoryHandler::Create]")
	return http.Created(nil, nil)
}

func NewInventoryHandler(logger interfaces.ILogger, validator interfaces.IValidator) IInventoryHandler {
	return inventoryHandler{
		logger:    logger,
		validator: validator,
	}
}
