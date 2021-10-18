package interfaces

import (
	"webapi/pkg/domain/dtos"
)

type ITokenManager interface {
	GenerateToken(tokenData dtos.TokenDataDto) (string, error)
	VerifyToken(token string) (*dtos.AuthenticatedUserDto, error)
}
