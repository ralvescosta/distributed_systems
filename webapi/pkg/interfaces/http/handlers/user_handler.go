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
	logger   interfaces.ILogger
	useCases usecases.ICreateUserUseCase
}

func (pst usersHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.CreateUserRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		return http.BadRequest("body is required", nil)
	}

	result, err := pst.useCases.CreateUser(httpRequest.Ctx, model.ToCreateUserDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Created(models.ToCreateUserResponse(result), nil)
}

func NewUsersHandler(logger interfaces.ILogger, useCases usecases.ICreateUserUseCase) IUsersHandler {
	return usersHandler{
		logger:   logger,
		useCases: useCases,
	}
}
