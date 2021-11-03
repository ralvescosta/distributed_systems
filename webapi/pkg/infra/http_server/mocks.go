package httpserver

import (
	"net/http"
	"net/http/httptest"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/infra/logger"

	"github.com/gin-gonic/gin"
)

type httpServerToTest struct {
	server    HttpServer
	loggerSpy interfaces.ILogger
}

func (pst httpServerToTest) doRequest(method, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}

	w := httptest.NewRecorder()
	pst.server.server.ServeHTTP(w, req)

	return w.Result(), nil
}

func newHttpServerToTest() httpServerToTest {
	gin.SetMode(gin.TestMode)

	loggerSpy := logger.NewLoggerSpy()

	return httpServerToTest{
		server:    HttpServer{logger: loggerSpy},
		loggerSpy: loggerSpy,
	}
}
