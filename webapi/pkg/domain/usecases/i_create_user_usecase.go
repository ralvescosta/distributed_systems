package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type ICreateUserUseCase interface {
	Perform(ctx context.Context, txn interface{}, dto dtos.CreateUserDto) (dtos.CreatedUserDto, error)
}
