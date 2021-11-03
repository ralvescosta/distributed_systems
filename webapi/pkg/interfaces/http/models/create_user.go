package models

import (
	"webapi/pkg/domain/dtos"
)

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
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
	ExpireIn    string `json:"expire_in"`
}

func ToCreateUserResponse(user dtos.CreatedUserDto) CreateUserResponse {
	return CreateUserResponse{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		AccessToken: user.AccessToken,
		Kind:        user.Kind,
		ExpireIn:    user.ExpireIn.String(),
	}
}
