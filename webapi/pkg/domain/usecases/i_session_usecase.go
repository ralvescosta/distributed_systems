package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type ISessionUseCase interface {
	Perform(ctx context.Context, dto dtos.SignInDto) (dtos.SessionDto, error)
}
