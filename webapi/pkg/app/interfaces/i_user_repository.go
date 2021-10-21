package interfaces

import (
	"context"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
)

type IUserRepository interface {
	FindById(ctx context.Context, txn interface{}, id int) (*entities.User, error)
	FindByEmail(ctx context.Context, txn interface{}, email string) (*entities.User, error)
	Create(ctx context.Context, txn interface{}, dto dtos.CreateUserDto) (*entities.User, error)
}
