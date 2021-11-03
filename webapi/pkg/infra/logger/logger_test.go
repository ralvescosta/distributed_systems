package logger

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Execute_GetHandleFunc_For_Test_Environment(t *testing.T) {
	sut := newLoggerToTest()

	handle := sut.logger.GetHandleFunc()

	ty := reflect.TypeOf(handle)

	assert.Equal(t, ty.Name(), "HandlerFunc", "Get Handle Func return wrong type to test environment")
}

func Test_Should_Execute_GetHandleFunc_For_Prod_Environment(t *testing.T) {
	sut := newLoggerToTest()
	os.Setenv("GO_ENV", "production")

	handle := sut.logger.GetHandleFunc()

	ty := reflect.TypeOf(handle)

	assert.Equal(t, ty.Name(), "HandlerFunc", "Get Handle Func return wrong type to test environment")
}

func Test_Should_Execute_ProductionLoggerFormater_Correctly(t *testing.T) {
	os.Setenv("GO_ENV", "production")
	sut := newLoggerToTest()

	req := &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(nil))),
		Header: http.Header{
			"op": []string{"op"},
		},
	}
	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())
	contextMock.Params = []gin.Param{{Key: "key", Value: "value"}}
	contextMock.Request = req

	param := reflect.ValueOf(contextMock)

	reflect.ValueOf(sut.logger).MethodByName("HttpRequestLogger").Call([]reflect.Value{param})
}

func Test_Should_Call_Logger_Methods_Correctly(t *testing.T) {
	sut := newLoggerToTest()

	sut.logger.Info("some message")
	sut.logger.Debug("some message")
	sut.logger.Warn("some message")
	sut.logger.Error("some message")
}
