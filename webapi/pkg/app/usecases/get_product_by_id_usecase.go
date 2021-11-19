package usecases

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type getProductByIdUseCase struct {
	inventoryClient interfaces.IIventoryClient
}

func (pst getProductByIdUseCase) Perform(ctx context.Context, id string) (dtos.ProductDto, error) {
	return pst.inventoryClient.GetProductById(ctx, id)
}

func NewGetProductByIdUseCase(inventoryClient interfaces.IIventoryClient) usecases.IGetProductByIdUseCase {
	return getProductByIdUseCase{inventoryClient}
}
