package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IAuthenticationUseCase interface {
	Perform(ctx context.Context, txn interface{}, accessToken string) (dtos.AuthenticatedUserDto, error)
}
