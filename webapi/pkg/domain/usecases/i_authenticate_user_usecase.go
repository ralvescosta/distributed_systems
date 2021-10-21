package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IAuthenticateUserUseCase interface {
	Perform(ctx context.Context, txn interface{}, dto dtos.AuthenticateUserDto) (dtos.AuthenticatedUserDto, error)
}
