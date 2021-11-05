package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IValidationTokenUseCase interface {
	Perform(ctx context.Context, txn interface{}, accessToken string) (dtos.SessionDto, error)
}
