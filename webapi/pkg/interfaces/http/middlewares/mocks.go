package middlewares

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type authMiddlewareToTest struct {
	middleware IAuthMiddleware
}

func newAuthMiddlewareToTest(useCaseError error) authMiddlewareToTest {
	middleware := NewAuthMiddleware(validateTokenUseCaseSpy{useCaseError})
	return authMiddlewareToTest{middleware}
}

type validateTokenUseCaseSpy struct {
	useCaseError error
}

func (pst validateTokenUseCaseSpy) Perform(ctx context.Context, txn interface{}, accessToken string) (dtos.SessionDto, error) {
	return dtos.SessionDto{Id: 1}, pst.useCaseError
}
