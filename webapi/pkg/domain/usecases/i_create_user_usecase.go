package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type ICreateUserUseCase interface {
	CreateUser(ctx context.Context, dto dtos.CreateUserDto) (dtos.CreatedUserDto, error)
}
