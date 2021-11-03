package handlers

import (
	"context"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
	"webapi/pkg/infra/logger"
)

type userHandlerToTest struct {
	handler   IUsersHandler
	loggerSpy interfaces.ILogger
	useCase   usecases.ICreateUserUseCase
}

func newUserHandlerToTest() userHandlerToTest {
	loggerSpy := logger.NewLoggerSpy()
	useCase := createUserUseCaseSpy{}
	validatorSpy := _validatorSpy{}
	handler := NewUsersHandler(loggerSpy, useCase, validatorSpy)

	return userHandlerToTest{handler, loggerSpy, useCase}
}

type _validatorSpy struct{}

func (_validatorSpy) ValidateStruct(m interface{}) []dtos.ValidatedDto {
	return nil
}

type createUserUseCaseSpy struct{}

func (createUserUseCaseSpy) Perform(ctx context.Context, txn interface{}, dto dtos.CreateUserDto) (dtos.CreatedUserDto, error) {
	return dtos.CreatedUserDto{}, nil
}
