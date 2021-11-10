package cmd

import (
	"fmt"
	"net/http"
	"webapi/pkg/infra/environments"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

func WebApi() error {
	if err := environments.Configure(); err != nil {
		return err
	}

	container := NewContainer()

	// Server setup
	container.httpServer.Setup()

	//middlewares
	container.httpServer.RegisterMiddleware(func(ctx *gin.Context) {
		span := opentracing.GlobalTracer().StartSpan(fmt.Sprintf("HTTP %s %s", ctx.Request.Method, ctx.Request.RequestURI))
		defer span.Finish()

		tracerCtx := opentracing.ContextWithSpan(ctx.Request.Context(), span)

		ctx.Set("tracerCtx", tracerCtx)
		ctx.Next()

		responseStatusCode := ctx.Writer.Status()
		span.SetTag("http.status_code", responseStatusCode)

		if responseStatusCode >= http.StatusBadRequest {
			span.SetTag("error", true)
		}
	})

	// Router register
	container.usersRoutes.Register(container.httpServer)
	container.authenticationRoutes.Register(container.httpServer)
	container.inventoryRoutes.Register(container.httpServer)

	if err := container.httpServer.Run(); err != nil {
		return err
	}
	return nil
}
