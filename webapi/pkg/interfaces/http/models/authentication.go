package models

import "webapi/pkg/domain/dtos"

type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (pst AuthenticationRequest) ToAuthenticateUserDto() dtos.AuthenticateUserDto {
	return dtos.AuthenticateUserDto{
		Email:    pst.Email,
		Password: pst.Password,
	}
}

type AuthenticationResponse struct {
	AccessToken string `json:"access_token"`
	Kind        string `json:"kind"`
	ExpiredIn   string `json:"expired_in"`
}

func ToAuthenticationResponse(dto dtos.AuthenticatedUserDto) AuthenticationResponse {
	return AuthenticationResponse{
		AccessToken: dto.AccessToken,
		Kind:        dto.Kind,
		ExpiredIn:   dto.ExpireIn.Format(("2006-01-02 15:04:05")),
	}
}
