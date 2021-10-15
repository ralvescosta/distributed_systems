package httpserver

import (
	"errors"
	"fmt"
	"os"
	"webapi/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
)

type IHttpServer interface {
	Setup()
	RegistreRoute(method HttpMethod, path string, handlers ...gin.HandlerFunc) error
	Run() error
}

type HttpServer struct {
	server *gin.Engine
	logger interfaces.ILogger
}

var httpServerWrapper = gin.New

func (hs *HttpServer) Setup() {
	hs.server = httpServerWrapper()
	hs.server.Use(hs.logger.GetHandleFunc())
}

func (hs HttpServer) RegistreRoute(method HttpMethod, path string, handlers ...gin.HandlerFunc) error {
	switch method {
	case HttpMethod("POST"):
		hs.server.POST(path, handlers...)

	case HttpMethod("GET"):
		hs.server.GET(path, handlers...)

	case HttpMethod("PUT"):
		hs.server.PUT(path, handlers...)

	case HttpMethod("DELETE"):
		hs.server.DELETE(path, handlers...)
	default:
		return errors.New("http method not allowed")
	}
	return nil
}

func (hs HttpServer) Run() error {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	return hs.server.Run(fmt.Sprintf("%s:%s", host, port))
}

func NewHttpServer(logger interfaces.ILogger) IHttpServer {
	return &HttpServer{
		logger: logger,
	}
}
