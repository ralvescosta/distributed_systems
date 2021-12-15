package handlers

import (
	"encoding/json"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IPurchaseHandler interface {
	Create(req http.HttpRequest) http.HttpResponse
}

type purchaseHandler struct {
	logger    interfaces.ILogger
	validator interfaces.IValidator
	usecase   usecases.IPurchaseUseCase
}

func (pst purchaseHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.CreateProductRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		pst.logger.Error(err.Error())
		return http.BadRequest(models.StringToErrorResponse("body is required"), nil)
	}

	if validationErrs := pst.validator.ValidateStruct(model); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return http.BadRequest(models.StringToErrorResponse(validationErrs[0].Message), nil)
	}

	if err := pst.usecase.Perform(httpRequest.Ctx, model.ToCreatePutchaseDto()); err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Ok(nil, nil)
}

func NewPurchaseHandler(logger interfaces.ILogger, validator interfaces.IValidator, usecase usecases.IPurchaseUseCase) IPurchaseHandler {
	return purchaseHandler{
		logger,
		validator,
		usecase,
	}
}
