package interfaces

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IIventoryClient interface {
	GetProductById(ctx context.Context, id string) (dtos.ProductDto, error)
	RegisterProduct(ctx context.Context, product dtos.ProductDto) (dtos.ProductDto, error)
}
