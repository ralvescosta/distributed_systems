package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type ICreateProductUseCase interface {
	Perform(ctx context.Context, dto dtos.ProductDto) (dtos.ProductDto, error)
}
