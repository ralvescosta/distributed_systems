package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type ISessionUseCase interface {
	Perform(ctx context.Context, txn interface{}, dto dtos.SignInDto) (dtos.SessionDto, error)
}
