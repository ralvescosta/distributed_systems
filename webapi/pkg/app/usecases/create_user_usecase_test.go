package usecases

import (
	"context"
	"errors"
	"testing"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/entities"

	"github.com/stretchr/testify/assert"

	inernalError "webapi/pkg/app/errors"
)

func Test_CreateUserUC_Should_Create_User_Correctly(t *testing.T) {
	configs := map[string]mockConfigure{
		"some": {},
	}
	sut := newCreateUserUseCaseToTest(configs)

	result, err := sut.useCase.Perform(context.Background(), dtos.CreateUserDto{})

	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func Test_CreateUserUC_Should_Return_InternalError_If_Some_Error_Occur_In_Repository(t *testing.T) {
	configs := map[string]mockConfigure{
		"userRepository": {
			method:       "FindByEmail",
			customError:  errors.New("some error"),
			customResult: &entities.User{},
		},
	}
	sut := newCreateUserUseCaseToTest(configs)

	_, err := sut.useCase.Perform(context.Background(), dtos.CreateUserDto{})

	assert.Error(t, err)
	assert.IsType(t, err, inernalError.InternalError{})
}

func Test_CreateUserUC_Should_Return_ConflictError_If_Email_Already_Exist(t *testing.T) {
	configs := map[string]mockConfigure{
		"userRepository": {
			method:       "FindByEmail",
			customError:  nil,
			customResult: &entities.User{},
		},
	}
	sut := newCreateUserUseCaseToTest(configs)

	_, err := sut.useCase.Perform(context.Background(), dtos.CreateUserDto{})

	assert.Error(t, err)
	assert.IsType(t, err, inernalError.ConflictError{})
}

func Test_CreateUserUC_Should_Return_InternalError_If_Some_Error_Occur_In_Hasher(t *testing.T) {
	configs := map[string]mockConfigure{
		"hasher": {
			method:       "Hahser",
			customError:  errors.New("some error"),
			customResult: "",
		},
	}
	sut := newCreateUserUseCaseToTest(configs)

	_, err := sut.useCase.Perform(context.Background(), dtos.CreateUserDto{})

	assert.Error(t, err)
	assert.IsType(t, err, inernalError.InternalError{})
}

func Test_CreateUserUC_Should_Return_InternalError_If_Some_Error_Occur_When_Create_User(t *testing.T) {
	configs := map[string]mockConfigure{
		"userRepository": {
			method:       "Create",
			customError:  errors.New("some error"),
			customResult: &entities.User{},
		},
	}
	sut := newCreateUserUseCaseToTest(configs)

	_, err := sut.useCase.Perform(context.Background(), dtos.CreateUserDto{})

	assert.Error(t, err)
	assert.IsType(t, err, inernalError.InternalError{})
}

func Test_CreateUserUC_Should_Return_InternalError_If_Some_Error_Occur_When_Create_Token(t *testing.T) {
	configs := map[string]mockConfigure{
		"tokenManager": {
			method:       "GenerateToken",
			customError:  errors.New("some error"),
			customResult: "",
		},
	}
	sut := newCreateUserUseCaseToTest(configs)

	_, err := sut.useCase.Perform(context.Background(), dtos.CreateUserDto{})

	assert.Error(t, err)
	assert.IsType(t, err, inernalError.InternalError{})
}
