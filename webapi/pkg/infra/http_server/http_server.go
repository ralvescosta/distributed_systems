package httpserver

import (
	"errors"
	"fmt"
	"os"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

type IHttpServer interface {
	Setup()
	RegistreRoute(method http.HttpMethod, path string, handlers ...gin.HandlerFunc) error
	RegisterMiddleware(middleware ...gin.HandlerFunc)
	Run() error
}

type HttpServer struct {
	server *gin.Engine
	logger interfaces.ILogger
}

var httpServerWrapper = gin.New

func (pst *HttpServer) Setup() {
	pst.server = httpServerWrapper()
	pst.server.Use(pst.logger.GetHandleFunc())
}

func (pst HttpServer) RegistreRoute(method http.HttpMethod, path string, handlers ...gin.HandlerFunc) error {
	switch method {
	case http.HttpMethod("POST"):
		pst.server.POST(path, handlers...)

	case http.HttpMethod("GET"):
		pst.server.GET(path, handlers...)

	case http.HttpMethod("PUT"):
		pst.server.PUT(path, handlers...)

	case http.HttpMethod("DELETE"):
		pst.server.DELETE(path, handlers...)
	default:
		return errors.New("http method not allowed")
	}
	return nil
}

func (pst HttpServer) RegisterMiddleware(middleware ...gin.HandlerFunc) {
	pst.server.Use(middleware...)
}

func (pst HttpServer) Run() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return pst.server.Run(fmt.Sprintf("%s:%s", host, port))
}

func NewHttpServer(logger interfaces.ILogger) IHttpServer {
	return &HttpServer{
		logger: logger,
	}
}
