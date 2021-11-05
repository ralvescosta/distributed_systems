package handlers

import (
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IInventoryHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
	GetById(httpRequest http.HttpRequest) http.HttpResponse
}

type inventoryHandler struct {
	logger    interfaces.ILogger
	validator interfaces.IValidator
	useCase   usecases.IGetBookByIdUseCase
}

func (pst inventoryHandler) GetById(httpRequest http.HttpRequest) http.HttpResponse {
	id, ok := httpRequest.Params["id"]
	if !ok {
		return http.BadRequest(models.StringToErrorResponse("id is required"), nil)
	}

	result, err := pst.useCase.Perform(httpRequest.Ctx, httpRequest.Txn, id)
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Created(models.ToGetByIdResponse(result), nil)
}

func (pst inventoryHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	pst.logger.Debug("[InventoryHandler::Create]")
	return http.Created(nil, nil)
}

func NewInventoryHandler(logger interfaces.ILogger, validator interfaces.IValidator, useCase usecases.IGetBookByIdUseCase) IInventoryHandler {
	return inventoryHandler{
		logger,
		validator,
		useCase,
	}
}
