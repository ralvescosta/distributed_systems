package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IAuthenticationUseCase interface {
	Perform(ctx context.Context, txn interface{}, dto dtos.AuthenticationDto) (dtos.AuthenticatedUserDto, error)
}
