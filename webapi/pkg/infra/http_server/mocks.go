package httpserver

import (
	"net/http"
	"net/http/httptest"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

type HttpServerToTest struct {
	server    HttpServer
	loggerSpy interfaces.ILogger
}

func (pst HttpServerToTest) doRequest(method, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	w := httptest.NewRecorder()
	pst.server.server.ServeHTTP(w, req)

	return w.Result(), nil
}

func NewHttpServerToTest() HttpServerToTest {
	gin.SetMode(gin.TestMode)

	loggerSpy := logger.NewLoggerSpy()

	return HttpServerToTest{
		server:    HttpServer{logger: loggerSpy},
		loggerSpy: loggerSpy,
	}
}
