package http

import (
	"errors"
	"net/http"
	"testing"
	internalErrors "webapi/pkg/app/errors"
	"webapi/pkg/interfaces/http/models"

	"github.com/stretchr/testify/assert"
)

type someBody struct{}

func Test_OkFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := Ok(someBody{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusOK)
	assert.IsType(t, result.Body, someBody{})
}

func Test_CreatedFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := Created(someBody{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusCreated)
	assert.IsType(t, result.Body, someBody{})
}

func Test_NoContentFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := NoContent(http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusNoContent)
	assert.IsType(t, result.Body, nil)
}

func Test_BadRequestFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := BadRequest(models.ErrorResponse{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusBadRequest)
	assert.IsType(t, result.Body, models.ErrorResponse{})
}

func Test_UnauthorizedFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := Unauthorized(models.ErrorResponse{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusUnauthorized)
	assert.IsType(t, result.Body, models.ErrorResponse{})
}

func Test_ForbidenFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := Forbiden(models.ErrorResponse{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusForbidden)
	assert.IsType(t, result.Body, models.ErrorResponse{})
}

func Test_NotFoundFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := NotFound(models.ErrorResponse{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusNotFound)
	assert.IsType(t, result.Body, models.ErrorResponse{})
}

func Test_ConflictFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := Conflict(models.ErrorResponse{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusConflict)
	assert.IsType(t, result.Body, models.ErrorResponse{})
}

func Test_InternalServerErrorFunc_Http_Should_Return_Ok_StatusCode(t *testing.T) {
	result := InternalServerError(models.ErrorResponse{}, http.Header{})

	assert.Equal(t, result.StatusCode, http.StatusInternalServerError)
	assert.IsType(t, result.Body, models.ErrorResponse{})
}

func Test_ErrorResponseMapper(t *testing.T) {
	type inputs struct {
		err    error
		status int
	}
	inputsToTest := []inputs{
		{err: internalErrors.NewBadRequestError(""), status: http.StatusBadRequest},
		{err: internalErrors.NewUnauthorizeError(""), status: http.StatusUnauthorized},
		{err: internalErrors.NewNotFoundError(""), status: http.StatusNotFound},
		{err: internalErrors.NewConflictError(""), status: http.StatusConflict},
		{err: errors.New(""), status: http.StatusInternalServerError},
	}

	for _, in := range inputsToTest {
		result := ErrorResponseMapper(in.err, http.Header{})

		assert.Equal(t, result.StatusCode, in.status)
	}
}
