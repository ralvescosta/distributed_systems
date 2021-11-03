package handlers

import (
	"encoding/json"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IUsersHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
}

type usersHandler struct {
	logger    interfaces.ILogger
	useCases  usecases.ICreateUserUseCase
	validator interfaces.IValidator
}

func (pst usersHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.CreateUserRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		return http.BadRequest(models.StringToErrorResponse("body is required"), nil)
	}

	if validationErrs := pst.validator.ValidateStruct(model); validationErrs != nil {
		pst.logger.Error(validationErrs[0].Message)
		return http.BadRequest(models.StringToErrorResponse(validationErrs[0].Message), nil)
	}

	result, err := pst.useCases.Perform(httpRequest.Ctx, httpRequest.Txn, model.ToCreateUserDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Created(models.ToCreateUserResponse(result), nil)
}

func NewUsersHandler(logger interfaces.ILogger, useCases usecases.ICreateUserUseCase, validator interfaces.IValidator) IUsersHandler {
	return usersHandler{
		logger,
		useCases,
		validator,
	}
}
