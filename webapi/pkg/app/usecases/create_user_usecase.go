package usecases

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
	"webapi/pkg/domain/usecases"
)

type createUserUseCase struct {
	repository interfaces.IUserRepository
}

func (pst createUserUseCase) CreateUser(ctx context.Context, dto dtos.CreateUserDto) (entities.User, error) {
	pst.repository.FindByEmail(ctx, dto.Email)
	return entities.User{}, nil
}

func NewCreateUserUseCase(repository interfaces.IUserRepository) usecases.ICreateUserUseCase {
	return createUserUseCase{
		repository: repository,
	}
}
