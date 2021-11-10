package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IValidationTokenUseCase interface {
	Perform(ctx context.Context, accessToken string) (dtos.SessionDto, error)
}
