package middlewares

import (
	"errors"
	"net/http"
	"testing"
	"webapi/pkg/domain/dtos"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Validate_Token_Correctly(t *testing.T) {
	sut := newAuthMiddlewareToTest(nil)

	result := sut.middleware.Perform(internalHttp.HttpRequest{
		Headers: http.Header{
			"Authorization": []string{"Bearer some_token"},
		},
	})

	assert.Equal(t, result.StatusCode, http.StatusOK)
	assert.IsType(t, result.Body, &dtos.SessionDto{})
}

func Test_Should_Return_Unauthorized_If_Has_No_Token(t *testing.T) {
	sut := newAuthMiddlewareToTest(nil)

	result := sut.middleware.Perform(internalHttp.HttpRequest{
		Headers: http.Header{},
	})

	assert.Equal(t, result.StatusCode, http.StatusUnauthorized)
}

func Test_Should_Return_Unauthorized_If_Token_Is_Unformatted(t *testing.T) {
	sut := newAuthMiddlewareToTest(nil)

	result := sut.middleware.Perform(internalHttp.HttpRequest{
		Headers: http.Header{
			"Authorization": []string{"some_token"},
		},
	})

	assert.Equal(t, result.StatusCode, http.StatusUnauthorized)
}

func Test_Should_Return_Unauthorized_If_Some_Error_Occur_In_UseCase(t *testing.T) {
	sut := newAuthMiddlewareToTest(errors.New("some error"))

	result := sut.middleware.Perform(internalHttp.HttpRequest{
		Headers: http.Header{
			"Authorization": []string{"Bearer some_token"},
		},
	})

	assert.Equal(t, result.StatusCode, http.StatusUnauthorized)
}
