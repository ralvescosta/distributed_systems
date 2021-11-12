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

type sessionUseCase struct {
	repository   interfaces.IUserRepository
	hasher       interfaces.IHasher
	tokenManager interfaces.ITokenManager
	logger       interfaces.ILogger
}

func (pst sessionUseCase) Perform(ctx context.Context, dto dtos.SignInDto) (dtos.SessionDto, error) {
	user, err := pst.repository.FindByEmail(ctx, dto.Email)
	if err != nil {
		return dtos.SessionDto{}, err
	}

	if user == nil {
		return dtos.SessionDto{}, errors.NewNotFoundError("Email not found")
	}

	if !pst.hasher.Verify(dto.Password, user.Password) {
		return dtos.SessionDto{}, errors.NewBadRequestError("Wrong password")
	}

	expireIn := time.Now().Add(time.Hour * 1)
	accessToken, err := pst.tokenManager.GenerateToken(dtos.TokenDataDto{
		Id:       user.Id,
		Audience: os.Getenv("APP_ISSUER"),
		ExpireIn: expireIn,
	})
	if err != nil {
		return dtos.SessionDto{}, err
	}

	return dtos.SessionDto{
		AccessToken: accessToken,
		Kind:        os.Getenv("TOKEN_KIND"),
		ExpireIn:    expireIn,
	}, nil

}

func NewSessionUseCase(repository interfaces.IUserRepository, hasher interfaces.IHasher, tokenManager interfaces.ITokenManager, logger interfaces.ILogger) usecases.ISessionUseCase {
	return sessionUseCase{
		repository,
		hasher,
		tokenManager,
		logger,
	}
}
