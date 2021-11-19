package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IGetProductByIdUseCase interface {
	Perform(ctx context.Context, id string) (dtos.ProductDto, error)
}
