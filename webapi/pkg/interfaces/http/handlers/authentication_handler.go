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
	logger   interfaces.ILogger
	useCases usecases.IAuthenticationUseCase
}

func (pst authenticationHandler) Create(httpRequest http.HttpRequest) http.HttpResponse {
	model := models.AuthenticationRequest{}
	if err := json.Unmarshal(httpRequest.Body, &model); err != nil {
		return http.BadRequest("body is required", nil)
	}

	result, err := pst.useCases.Perform(httpRequest.Ctx, model.ToAuthenticationDto())
	if err != nil {
		return http.ErrorResponseMapper(err, nil)
	}

	return http.Ok(models.ToAuthenticationResponse(result), nil)
}

func NewAuthenticationHandler(logger interfaces.ILogger, useCases usecases.IAuthenticationUseCase) IAuthenticationHandler {
	return authenticationHandler{
		logger:   logger,
		useCases: useCases,
	}
}
