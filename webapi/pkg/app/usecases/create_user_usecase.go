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

type createUserUseCase struct {
	repository   interfaces.IUserRepository
	hasher       interfaces.IHasher
	tokenManager interfaces.ITokenManager
}

func (pst createUserUseCase) Perform(ctx context.Context, txn interface{}, dto dtos.CreateUserDto) (dtos.CreatedUserDto, error) {
	user, err := pst.repository.FindByEmail(ctx, txn, dto.Email)
	if err != nil {
		return dtos.CreatedUserDto{}, errors.NewInternalError(err.Error())
	}

	if user != nil {
		return dtos.CreatedUserDto{}, errors.NewConflictError("Email already registered")
	}

	hashedPassword, err := pst.hasher.Hahser(dto.Password)
	if err != nil {
		return dtos.CreatedUserDto{}, errors.NewInternalError(err.Error())
	}

	dto.Password = hashedPassword
	user, err = pst.repository.Create(ctx, txn, dto)
	if err != nil {
		return dtos.CreatedUserDto{}, errors.NewInternalError(err.Error())
	}

	expireIn := time.Now().Add(time.Hour)
	accessToken, err := pst.tokenManager.GenerateToken(dtos.TokenDataDto{
		Id:       user.Id,
		Audience: "WebApi",
		ExpireIn: expireIn,
	})
	if err != nil {
		return dtos.CreatedUserDto{}, errors.NewInternalError(err.Error())
	}

	return dtos.CreatedUserDto{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
		AccessToken: accessToken,
		Kind:        os.Getenv("TOKEN_KIND"),
		ExpireIn:    expireIn,
	}, nil
}

func NewCreateUserUseCase(repository interfaces.IUserRepository, hasher interfaces.IHasher, tokenManager interfaces.ITokenManager) usecases.ICreateUserUseCase {
	return createUserUseCase{
		repository,
		hasher,
		tokenManager,
	}
}
