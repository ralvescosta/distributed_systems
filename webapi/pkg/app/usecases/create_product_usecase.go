package usecases

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type createProductUseCase struct {
	inventoryClient interfaces.IIventoryClient
}

func (pst createProductUseCase) Perform(ctx context.Context, dto dtos.ProductDto) (dtos.ProductDto, error) {
	return pst.inventoryClient.RegisterProduct(ctx, dto)
}

func NewCreateProductUseCase(inventoryClient interfaces.IIventoryClient) usecases.ICreateProductUseCase {
	return createProductUseCase{
		inventoryClient,
	}
}
