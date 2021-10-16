package usecases

import (
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
	"webapi/pkg/domain/usecases"
)

type createUserUseCase struct{}

func (createUserUseCase) CreateUser(dto dtos.CreateUserDto) (entities.User, error) {
	return entities.User{}, nil
}

func NewCreateUserUseCase() usecases.ICreateUserUseCase {
	return createUserUseCase{}
}
