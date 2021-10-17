package interfaces

import (
	"context"
	"webapi/pkg/domain/entities"
)

type IUserRepository interface {
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	Create(ctx context.Context) (*entities.User, error)
}
