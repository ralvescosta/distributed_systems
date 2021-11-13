package usecases

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type getBookByIdUseCase struct {
	inventoryClient interfaces.IIventoryClient
}

func (pst getBookByIdUseCase) Perform(ctx context.Context, id string) (dtos.BookDto, error) {
	pst.inventoryClient.GetProductById(ctx, id)

	return dtos.BookDto{}, nil
}

func NewGetBookByIdUseCase(inventoryClient interfaces.IIventoryClient) usecases.IGetBookByIdUseCase {
	return getBookByIdUseCase{inventoryClient}
}
