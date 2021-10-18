package usecases

import (
	"context"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
	"webapi/pkg/domain/usecases"
)

type createUserUseCase struct {
	repository interfaces.IUserRepository
	hasher     interfaces.IHasher
}

func (pst createUserUseCase) CreateUser(ctx context.Context, dto dtos.CreateUserDto) (entities.User, error) {
	user, err := pst.repository.FindByEmail(ctx, dto.Email)
	if err != nil {
		return entities.User{}, errors.NewInternalError(err.Error())
	}

	if user != nil {
		return entities.User{}, errors.NewConflictError("Email already registered")
	}

	hashedPassword, err := pst.hasher.Hahser(dto.Password)
	if err != nil {
		return entities.User{}, errors.NewInternalError(err.Error())
	}

	user.Password = hashedPassword
	user, err = pst.repository.Create(ctx, dto)
	if err != nil {
		return entities.User{}, errors.NewInternalError(err.Error())
	}

	return *user, nil
}

func NewCreateUserUseCase(repository interfaces.IUserRepository, hasher interfaces.IHasher) usecases.ICreateUserUseCase {
	return createUserUseCase{
		repository: repository,
		hasher:     hasher,
	}
}
