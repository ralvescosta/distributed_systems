package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"webapi/pkg/interfaces/http"
)

func Test_Should_Build_Http_Request_Correctly(t *testing.T) {
	ginCtx := createMockedGinContext(createMockedHttpRequest(false))
	request, err := GetHttpRequest(ginCtx)

	assert.NoError(t, err)
	assert.IsType(t, request, http.HttpRequest{})
}

func Test_Should_Return_Err_If_Some_Error_Occur_In_Body_Reader(t *testing.T) {
	ginCtx := createMockedGinContext(createMockedHttpRequest(true))
	request, err := GetHttpRequest(ginCtx)

	assert.Error(t, err)
	assert.Nil(t, request.Body)
}
