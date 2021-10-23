package handlers

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IAuthenticationHandler interface {
	Create(httpRequest http.HttpRequest) http.HttpResponse
}

type authenticationHandler struct {
	logger   interfaces.ILogger
	useCases usecases.IAuthenticateUserUseCase
}

func (pst authenticationHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.AuthenticationRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		return http.BadRequest("body is required", nil)
	}

	v := validator.New()
	err := v.Struct(model)
	if err != nil {
		pst.logger.Info(err.Error())
	}

	result, err := pst.useCases.Perform(httpRequest.Ctx, httpRequest.Txn, model.ToAuthenticateUserDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Ok(models.ToAuthenticationResponse(result), nil)
}

func NewAuthenticationHandler(logger interfaces.ILogger, useCases usecases.IAuthenticateUserUseCase) IAuthenticationHandler {
	return authenticationHandler{
		logger:   logger,
		useCases: useCases,
	}
}
