package models

import "webapi/pkg/domain/dtos"

type PurchaseProduct struct {
	Id     string `json:"id" validate:"required"`
	Number uint   `json:"number" validate:"required"`
	Amount uint   `json:"amount" validate:"required"`
}

type CreateProductRequest struct {
	OrderId     string            `json:"order_id" validate:"required,uuid4"`
	Products    []PurchaseProduct `json:"products" validate:"required"`
	UserId      int               `json:"user_id" validate:"required"`
	PurchasedAt string            `json:"purchased_at" validate:"required"`
}

func (pst CreateProductRequest) ToCreatePutchaseDto() dtos.CreatePurchaseDto {
	return dtos.CreatePurchaseDto{
		OrderId: pst.OrderId,
	}
}
