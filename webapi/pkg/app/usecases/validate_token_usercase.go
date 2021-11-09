package usecases

import (
	"context"
	"webapi/pkg/app/errors"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type validatinTokenUseCase struct {
	repository   interfaces.IUserRepository
	tokenManager interfaces.ITokenManager
}

func (pst validatinTokenUseCase) Perform(ctx context.Context, txn interface{}, accessToken string) (dtos.SessionDto, error) {
	authenticatedUser, err := pst.tokenManager.VerifyToken(accessToken)
	if err != nil {
		return dtos.SessionDto{}, errors.NewInternalError("Some error occur whiling validate the access token")
	}

	user, err := pst.repository.FindById(ctx, txn, authenticatedUser.Id)
	if err != nil {
		return dtos.SessionDto{}, errors.NewInternalError("Some error occur whiling validate the access token")
	}
	if user == nil {
		return dtos.SessionDto{}, errors.NewUnauthorizeError("User no longer existe")
	}

	return *authenticatedUser, nil
}

func NewValidatinTokenUseCase(repository interfaces.IUserRepository, tokenManager interfaces.ITokenManager) usecases.IValidationTokenUseCase {
	return validatinTokenUseCase{
		repository,
		tokenManager,
	}
}
