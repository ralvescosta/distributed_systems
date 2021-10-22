package adapters

import (
	"bytes"
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

func CreateMockedHttpRequest(readerWithError bool) *http.Request {
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

func CreateMockedGinContext(req *http.Request) *gin.Context {
	contextMock, _ := gin.CreateTestContext(httptest.NewRecorder())
	contextMock.Params = []gin.Param{{Key: "key", Value: "value"}}
	contextMock.Request = req

	return contextMock
}

type HandlerAdaptToTest struct {
	adapt              gin.HandlerFunc
	loggerMock         interfaces.ILogger
	handlerCalledTimes *int
	handlerMock        func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse
	request            *http.Request
	ctx                *gin.Context
}

func NewHandlerAdaptToTest(readerWithError bool) HandlerAdaptToTest {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse {
		handlerCalledTimes++
		return internalHttp.HttpResponse{}
	}

	loggerMock := logger.NewLoggerSpy()
	req := CreateMockedHttpRequest(readerWithError)
	sut := HandlerAdapt(handlerMock, loggerMock)

	return HandlerAdaptToTest{
		handlerMock:        handlerMock,
		handlerCalledTimes: &handlerCalledTimes,
		loggerMock:         loggerMock,
		request:            req,
		adapt:              sut,
		ctx:                CreateMockedGinContext(req),
	}
}

type MiddlewareAdaptToTest struct {
	adapt              gin.HandlerFunc
	loggerMock         interfaces.ILogger
	handlerCalledTimes *int
	handlerMock        func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse
	request            *http.Request
	ctx                *gin.Context
}

func NewMiddlewareAdaptToTest(readerWithError bool) MiddlewareAdaptToTest {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse {
		handlerCalledTimes++
		return internalHttp.HttpResponse{}
	}

	loggerMock := logger.NewLoggerSpy()
	req := CreateMockedHttpRequest(readerWithError)
	sut := MiddlewareAdapt(handlerMock, loggerMock)

	return MiddlewareAdaptToTest{
		handlerMock:        handlerMock,
		handlerCalledTimes: &handlerCalledTimes,
		loggerMock:         loggerMock,
		request:            req,
		adapt:              sut,
		ctx:                CreateMockedGinContext(req),
	}
}
