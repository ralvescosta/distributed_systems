package usecases

import (
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
)

type ICreateUserUseCase interface {
	CreateUser(dto dtos.CreateUserDto) (entities.User, error)
}
