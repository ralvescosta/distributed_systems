package usecases

import (
	"context"
	"errors"
	"testing"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"

	internalError "webapi/pkg/app/errors"

	"github.com/stretchr/testify/assert"
)

func Test_ValidateTokenUC_Should_Validate_Token_Correctly(t *testing.T) {
	config := map[string]mockConfigure{
		"some": {},
	}

	sut := newValidationTokenUseCaseToTest(config)

	result, err := sut.useCase.Perform(context.Background(), nil, "some token")

	assert.NoError(t, err)
	assert.IsType(t, result, dtos.SessionDto{})
}

func Test_ValidateTokenUC_Should_Return_InternalError_If_Some_Error_Occur_In_VerifyToken(t *testing.T) {
	config := map[string]mockConfigure{
		"tokenManager": {
			method:       "VerifyToken",
			customResult: &dtos.SessionDto{},
			customError:  errors.New("some error"),
		},
	}

	sut := newValidationTokenUseCaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, "some token")

	assert.Error(t, err)
	assert.IsType(t, err, internalError.InternalError{})
}

func Test_ValidateTokenUC_Should_Return_InternalError_If_Some_Error_Occur_In_Repository(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindById",
			customResult: &entities.User{},
			customError:  errors.New("some error"),
		},
	}

	sut := newValidationTokenUseCaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, "some token")

	assert.Error(t, err)
	assert.IsType(t, err, internalError.InternalError{})
}

func Test_ValidateTokenUC_Should_Return_Unauthorized_If_User_Not_Found(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindById",
			customResult: nil,
			customError:  nil,
		},
	}

	sut := newValidationTokenUseCaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, "some token")

	assert.Error(t, err)
	assert.IsType(t, err, internalError.UnauthorizeError{})
}
