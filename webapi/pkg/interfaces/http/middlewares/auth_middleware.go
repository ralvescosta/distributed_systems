package middlewares

import (
	"strings"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
	"webapi/pkg/interfaces/http/models"
)

type IAuthMiddleware interface {
	Perform(httpRequest http.HttpRequest) http.HttpResponse
}

type authMiddleware struct {
	usecase usecases.IAuthenticationUseCase
}

func (pst authMiddleware) Perform(httpRequest http.HttpRequest) http.HttpResponse {
	authHeader := httpRequest.Headers.Get("Authorization")

	token := strings.Split(authHeader, " ")
	if token[0] != "Bearer" || len(token) < 2 {
		return http.Unauthorized(models.StringToErrorResponse("Authorization header unformatted"), httpRequest.Headers)
	}

	authenticatedUser, err := pst.usecase.Perform(httpRequest.Ctx, httpRequest.Txn, token[1])
	if err != nil {
		return http.Unauthorized(models.StringToErrorResponse("Invalid token"), httpRequest.Headers)
	}

	return http.Ok(&authenticatedUser, httpRequest.Headers)
}

func NewAuthMiddleware(usecase usecases.IAuthenticationUseCase) IAuthMiddleware {
	return authMiddleware{
		usecase,
	}
}
