package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IAuthenticationUseCase interface {
	Perform(ctx context.Context, dto dtos.AuthenticationDto) (dtos.AuthenticatedUserDto, error)
}
