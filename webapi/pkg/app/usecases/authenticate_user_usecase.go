package usecases

import (
	"context"
	"os"
	"time"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type authenticateUserUsecase struct {
	repository   interfaces.IUserRepository
	hasher       interfaces.IHasher
	tokenManager interfaces.ITokenManager
}

func (pst authenticateUserUsecase) Perform(ctx context.Context, txn interface{}, dto dtos.AuthenticateUserDto) (dtos.AuthenticatedUserDto, error) {
	user, err := pst.repository.FindByEmail(ctx, txn, dto.Email)
	if err != nil {
		return dtos.AuthenticatedUserDto{}, err
	}

	if user == nil {
		return dtos.AuthenticatedUserDto{}, errors.NewNotFoundError("Email not found")
	}

	if !pst.hasher.Verify(dto.Password, user.Password) {
		return dtos.AuthenticatedUserDto{}, errors.NewBadRequestError("Wrong password")
	}

	expireIn := time.Now().Add(time.Hour * 1)
	accessToken, err := pst.tokenManager.GenerateToken(dtos.TokenDataDto{
		Id:       user.Id,
		Audience: os.Getenv("APP_ISSUER"),
		ExpireIn: expireIn,
	})
	if err != nil {
		return dtos.AuthenticatedUserDto{}, err
	}

	return dtos.AuthenticatedUserDto{
		AccessToken: accessToken,
		Kind:        os.Getenv("TOKEN_KIND"),
		ExpireIn:    expireIn,
	}, nil

}

func NewAuthenticateUserUseCase(repository interfaces.IUserRepository, hasher interfaces.IHasher, tokenManager interfaces.ITokenManager) usecases.IAuthenticateUserUseCase {
	return authenticateUserUsecase{
		repository:   repository,
		hasher:       hasher,
		tokenManager: tokenManager,
	}
}
