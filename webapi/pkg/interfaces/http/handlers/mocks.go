package handlers

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/infra/logger"
	"webapi/pkg/interfaces/http/models"
)

type userHandlerToTest struct {
	handler    IUsersHandler
	loggerSpy  interfaces.ILogger
	useCase    usecases.ICreateUserUseCase
	mockedUser models.CreateUserRequest
}

func newUserHandlerToTest(validationFailure bool, useCaseError error) userHandlerToTest {
	loggerSpy := logger.NewLoggerSpy()
	useCase := createUserUseCaseSpy{useCaseError}
	validatorSpy := _validatorSpy{validationFailure}
	handler := NewUsersHandler(loggerSpy, useCase, validatorSpy)

	mockedUser := models.CreateUserRequest{
		Name:     "Some Name",
		Email:    "some@email.com",
		Password: "1234567",
	}

	return userHandlerToTest{handler, loggerSpy, useCase, mockedUser}
}

type _validatorSpy struct {
	validationFailure bool
}

func (pst _validatorSpy) ValidateStruct(m interface{}) []dtos.ValidatedDto {
	if pst.validationFailure {
		return []dtos.ValidatedDto{
			{
				IsValid: false,
			},
		}
	}
	return nil
}

type createUserUseCaseSpy struct {
	useCaseError error
}

func (pst createUserUseCaseSpy) Perform(ctx context.Context, txn interface{}, dto dtos.CreateUserDto) (dtos.CreatedUserDto, error) {
	return dtos.CreatedUserDto{}, pst.useCaseError
}

type sessionHandlerToTest struct {
	handler    ISessionHandler
	loggerSpy  interfaces.ILogger
	useCase    usecases.ISessionUseCase
	mockedUser models.CreateUserRequest
}

func newSessionHandlerToTest(validationFailure bool, useCaseError error) sessionHandlerToTest {
	loggerSpy := logger.NewLoggerSpy()
	useCase := sessionUseCaseSpy{useCaseError}
	validatorSpy := _validatorSpy{validationFailure}
	handler := NewSessionHandler(loggerSpy, useCase, validatorSpy)

	mockedUser := models.CreateUserRequest{
		Name:     "Some Name",
		Email:    "some@email.com",
		Password: "1234567",
	}

	return sessionHandlerToTest{handler, loggerSpy, useCase, mockedUser}
}

type sessionUseCaseSpy struct {
	useCaseError error
}

func (pst sessionUseCaseSpy) Perform(ctx context.Context, txn interface{}, dto dtos.SignInDto) (dtos.SessionDto, error) {
	return dtos.SessionDto{}, pst.useCaseError
}
