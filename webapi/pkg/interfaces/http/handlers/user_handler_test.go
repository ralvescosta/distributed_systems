package handlers

import (
	"encoding/json"
	"net/http"
	"testing"
	"webapi/pkg/app/errors"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Execute_Create_Correctly(t *testing.T) {
	sut := newUserHandlerToTest(false, nil)
	body, _ := json.Marshal(sut.mockedUser)

	result := sut.handler.Create(
		internalHttp.HttpRequest{
			Body: body,
		},
	)

	assert.Equal(t, result.StatusCode, http.StatusCreated)
}

func Test_Should_Returns_BadRequest_If_Has_No_Body(t *testing.T) {
	sut := newUserHandlerToTest(false, nil)

	result := sut.handler.Create(
		internalHttp.HttpRequest{},
	)

	assert.Equal(t, result.StatusCode, http.StatusBadRequest)
}

func Test_Should_Returns_BadRequest_If_There_Is_Validation_Error_In_Body(t *testing.T) {
	sut := newUserHandlerToTest(true, nil)
	body, _ := json.Marshal(sut.mockedUser)

	result := sut.handler.Create(
		internalHttp.HttpRequest{
			Body: body,
		},
	)

	assert.Equal(t, result.StatusCode, http.StatusBadRequest)
}

func Test_Should_Returns_Http4xx_If_Some_Error_Occur_In_UseCase(t *testing.T) {
	sut := newUserHandlerToTest(false, errors.NewConflictError("conflict"))
	body, _ := json.Marshal(sut.mockedUser)

	result := sut.handler.Create(
		internalHttp.HttpRequest{
			Body: body,
		},
	)

	assert.Equal(t, result.StatusCode, http.StatusConflict)
}
