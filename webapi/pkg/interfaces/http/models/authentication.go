package models

import "webapi/pkg/domain/dtos"

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (pst AuthenticationRequest) ToAuthenticationDto() dtos.AuthenticationDto {
	return dtos.AuthenticationDto{
		Email:    pst.Email,
		Password: pst.Password,
	}
}

type AuthenticationResponse struct {
	AccessToken string `json:"access_token"`
	Kind        string `json:"kind"`
	ExpiredAt   string `json:"expired_at"`
}

func ToAuthenticationResponse() AuthenticationResponse {
	return AuthenticationResponse{}
}
