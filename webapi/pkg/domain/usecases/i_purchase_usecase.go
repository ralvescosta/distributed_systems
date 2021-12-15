package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IPurchaseUseCase interface {
	Perform(ctx context.Context, dto dtos.CreatePurchaseDto) error
}
