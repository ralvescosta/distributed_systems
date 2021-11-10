package adapters

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/logger"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

type customReader struct{}

func (customReader) Read(p []byte) (int, error) {
	return 0, errors.New("some error")
}

func createMockedHttpRequest(readerWithError bool) *http.Request {
	var reader io.ReadCloser
	if readerWithError {
		reader = ioutil.NopCloser(customReader{})
	} else {
		reader = ioutil.NopCloser(bytes.NewBuffer([]byte(nil)))
	}

	return &http.Request{
		Body: reader,
		Header: http.Header{
			"op": []string{"op"},
		},
	}
}

func createMockedGinContext(req *http.Request) *gin.Context {
	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())
	contextMock.Params = []gin.Param{{Key: "key", Value: "value"}}
	contextMock.Request = req
	contextMock.Set("tracerCtx", context.Background())
	return contextMock
}

type handlerAdaptToTest struct {
	adapt              gin.HandlerFunc
	loggerMock         interfaces.ILogger
	handlerCalledTimes *int
	handlerMock        func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse
	request            *http.Request
	ctx                *gin.Context
}

func newHandlerAdaptToTest(readerWithError bool, httpStatusCodeResponse int) handlerAdaptToTest {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse {
		handlerCalledTimes++
		return internalHttp.HttpResponse{
			StatusCode: httpStatusCodeResponse,
		}
	}

	loggerMock := logger.NewLoggerSpy()
	req := createMockedHttpRequest(readerWithError)
	sut := HandlerAdapt(handlerMock, loggerMock)

	return handlerAdaptToTest{
		handlerMock:        handlerMock,
		handlerCalledTimes: &handlerCalledTimes,
		loggerMock:         loggerMock,
		request:            req,
		adapt:              sut,
		ctx:                createMockedGinContext(req),
	}
}

type middlewareAdaptToTest struct {
	adapt              gin.HandlerFunc
	loggerMock         interfaces.ILogger
	handlerCalledTimes *int
	handlerMock        func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse
	request            *http.Request
	ctx                *gin.Context
}

func newMiddlewareAdaptToTest(readerWithError bool, httpStatusCodeResponse int) middlewareAdaptToTest {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse {
		handlerCalledTimes++
		return internalHttp.HttpResponse{
			StatusCode: httpStatusCodeResponse,
		}
	}

	loggerMock := logger.NewLoggerSpy()
	req := createMockedHttpRequest(readerWithError)
	sut := MiddlewareAdapt(handlerMock, loggerMock)

	return middlewareAdaptToTest{
		handlerMock:        handlerMock,
		handlerCalledTimes: &handlerCalledTimes,
		loggerMock:         loggerMock,
		request:            req,
		adapt:              sut,
		ctx:                createMockedGinContext(req),
	}
}
