package usecases

import (
	"context"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type authenticationUseCase struct {
	repository   interfaces.IUserRepository
	tokenManager interfaces.ITokenManager
}

func (pst authenticationUseCase) Perform(ctx context.Context, txn interface{}, accessToken string) (dtos.AuthenticatedUserDto, error) {
	authenticatedUser, err := pst.tokenManager.VerifyToken(accessToken)
	if err != nil {
		return dtos.AuthenticatedUserDto{}, errors.NewInternalError("Some error occur whiling validate the access token")
	}

	user, err := pst.repository.FindById(ctx, txn, authenticatedUser.Id)
	if err != nil {
		return dtos.AuthenticatedUserDto{}, errors.NewInternalError("Some error occur whiling validate the access token")
	}
	if user == nil {
		return dtos.AuthenticatedUserDto{}, errors.NewUnauthorizeError("User no longer existe")
	}

	return *authenticatedUser, nil
}

func NewAuthenticationUseCase(repository interfaces.IUserRepository, tokenManager interfaces.ITokenManager) usecases.IAuthenticationUseCase {
	return authenticationUseCase{
		repository,
		tokenManager,
	}
}
