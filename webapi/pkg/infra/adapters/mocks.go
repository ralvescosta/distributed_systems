package adapters

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/logger"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

func CreateMockedHttpRequest() *http.Request {
	return &http.Request{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte(nil))),
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

func NewHandlerAdaptToTest() HandlerAdaptToTest {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse {
		handlerCalledTimes++
		return internalHttp.HttpResponse{}
	}

	loggerMock := logger.NewLoggerSpy()
	req := CreateMockedHttpRequest()
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

func NewMiddlewareAdaptToTest() HandlerAdaptToTest {
	handlerCalledTimes := 0
	handlerMock := func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse {
		handlerCalledTimes++
		return internalHttp.HttpResponse{}
	}

	loggerMock := logger.NewLoggerSpy()
	req := CreateMockedHttpRequest()
	sut := MiddlewareAdapt(handlerMock, loggerMock)

	return HandlerAdaptToTest{
		handlerMock:        handlerMock,
		handlerCalledTimes: &handlerCalledTimes,
		loggerMock:         loggerMock,
		request:            req,
		adapt:              sut,
		ctx:                CreateMockedGinContext(req),
	}
}
