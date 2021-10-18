package models

import (
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"
)

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (pst CreateUserRequest) ToCreateUserDto() dtos.CreateUserDto {
	return dtos.CreateUserDto{
		Name:     pst.Name,
		Email:    pst.Email,
		Password: pst.Password,
	}
}

type CreateUserResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
	Kind        string `json:"kind"`
	ExpiredAt   string `json:"expired_at"`
}

func ToCreateUserResponse(user entities.User) CreateUserResponse {
	return CreateUserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
