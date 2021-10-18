package interfaces

import (
	"context"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
)

type IUserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context, dto dtos.CreateUserDto) (*entities.User, error)
}
