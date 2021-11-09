package usecases

import (
	"context"
	"errors"
	"testing"
	internalError "webapi/pkg/app/errors"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"

	"github.com/stretchr/testify/assert"
)

func Test_SessionUC_Should_Validate_Token_Correctly(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindByEmail",
			customResult: &entities.User{},
			customError:  nil,
		},
	}

	sut := newSessionUsecaseToTest(config)

	result, err := sut.useCase.Perform(context.Background(), nil, dtos.SignInDto{})

	assert.NoError(t, err)
	assert.IsType(t, result, dtos.SessionDto{})
}

func Test_SessionUC_Should_Return_Error_If_Some_Error_Occur_In_Repository(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindById",
			customResult: nil,
			customError:  errors.New("some error"),
		},
	}

	sut := newSessionUsecaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, dtos.SignInDto{})

	assert.Error(t, err)
}

func Test_SessionUC_Should_Return_NotFoundError_If_Email_Not_Found(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindById",
			customResult: nil,
			customError:  nil,
		},
	}

	sut := newSessionUsecaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, dtos.SignInDto{})

	assert.Error(t, err)
	assert.IsType(t, err, internalError.NotFoundError{})
}

func Test_SessionUC_Should_Return_BadRequestError_If_Password_Is_Wrong(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindByEmail",
			customResult: &entities.User{},
			customError:  nil,
		},
		"hasher": {
			method:       "Verify",
			customResult: false,
			customError:  nil,
		},
	}

	sut := newSessionUsecaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, dtos.SignInDto{})

	assert.Error(t, err)
	assert.IsType(t, err, internalError.BadRequestError{})
}

func Test_SessionUC_Should_Return_Error_If_GenerateToken_Return_Error(t *testing.T) {
	config := map[string]mockConfigure{
		"userRepository": {
			method:       "FindByEmail",
			customResult: &entities.User{},
			customError:  nil,
		},
		"tokenManager": {
			method:       "GenerateToken",
			customResult: "",
			customError:  errors.New("some error"),
		},
	}

	sut := newSessionUsecaseToTest(config)

	_, err := sut.useCase.Perform(context.Background(), nil, dtos.SignInDto{})

	assert.Error(t, err)
}
