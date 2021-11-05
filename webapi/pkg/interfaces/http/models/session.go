package models

import "webapi/pkg/domain/dtos"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (pst SignInRequest) ToSignInDto() dtos.SignInDto {
	return dtos.SignInDto{
		Email:    pst.Email,
		Password: pst.Password,
	}
}

type SessionResponse struct {
	AccessToken string `json:"access_token"`
	Kind        string `json:"kind"`
	ExpiredIn   string `json:"expired_in"`
}

func ToSessionResponse(dto dtos.SessionDto) SessionResponse {
	return SessionResponse{
		AccessToken: dto.AccessToken,
		Kind:        dto.Kind,
		ExpiredIn:   dto.ExpireIn.Format(("2006-01-02 15:04:05")),
	}
}
