package handlers

import (
	"encoding/json"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type ISessionHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
}

type sessionHandler struct {
	logger    interfaces.ILogger
	useCases  usecases.ISessionUseCase
	validator interfaces.IValidator
}

func (pst sessionHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.SignInRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		pst.logger.Error(err.Error())
		return http.BadRequest(models.StringToErrorResponse("body is required"), nil)
	}

	if validationErrs := pst.validator.ValidateStruct(model); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return http.BadRequest(models.StringToErrorResponse(validationErrs[0].Message), nil)
	}

	result, err := pst.useCases.Perform(httpRequest.Ctx, httpRequest.Txn, model.ToSignInDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Ok(models.ToSessionResponse(result), nil)
}

func NewSessionHandler(logger interfaces.ILogger, useCases usecases.ISessionUseCase, validator interfaces.IValidator) ISessionHandler {
	return sessionHandler{
		logger,
		useCases,
		validator,
	}
}
