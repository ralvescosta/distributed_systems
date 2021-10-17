package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
)

type ICreateUserUseCase interface {
	CreateUser(ctx context.Context, dto dtos.CreateUserDto) (entities.User, error)
}
