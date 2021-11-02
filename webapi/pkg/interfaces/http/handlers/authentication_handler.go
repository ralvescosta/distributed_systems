package handlers

import (
	"encoding/json"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IAuthenticationHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
}

type authenticationHandler struct {
	logger    interfaces.ILogger
	useCases  usecases.IAuthenticateUserUseCase
	validator interfaces.IValidator
}

func (pst authenticationHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.AuthenticationRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		pst.logger.Error(err.Error())
		return http.BadRequest(models.StringToErrorResponse("body is required"), nil)
	}

	if validationErrs := pst.validator.ValidateStruct(model); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return http.BadRequest(models.StringToErrorResponse(validationErrs[0].Message), nil)
	}

	result, err := pst.useCases.Perform(httpRequest.Ctx, httpRequest.Txn, model.ToAuthenticateUserDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Ok(models.ToAuthenticationResponse(result), nil)
}

func NewAuthenticationHandler(logger interfaces.ILogger, useCases usecases.IAuthenticateUserUseCase, validator interfaces.IValidator) IAuthenticationHandler {
	return authenticationHandler{
		logger,
		useCases,
		validator,
	}
}
