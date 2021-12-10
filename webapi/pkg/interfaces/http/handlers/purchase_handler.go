package handlers

import (
	"webapi/pkg/domain/usecases"
	"webapi/pkg/interfaces/http"
)

type IPurchaseHandler interface {
	Create(req http.HttpRequest) http.HttpResponse
}

type purchaseHandler struct {
	usecase usecases.IPurchaseUseCase
}

func (purchaseHandler) Create(req http.HttpRequest) http.HttpResponse {
	return http.Created(nil, nil)
}

func NewPurchaseHandler(usecase usecases.IPurchaseUseCase) IPurchaseHandler {
	return purchaseHandler{usecase}
}
