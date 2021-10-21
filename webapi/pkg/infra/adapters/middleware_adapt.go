package adapters

import (
	"net/http"
	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

func MiddlewareAdapt(handler func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse, logger interfaces.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := GetHttpRequest(ctx)
		if err != nil {
			logger.Error("error while read request bytes")
			ctx.JSON(http.StatusInternalServerError, gin.H{})
		}

		result := handler(request)

		if result.StatusCode >= http.StatusBadRequest {
			ctx.JSON(result.StatusCode, result.Body)
			return
		}

		authenticatedUserDto, ok := result.Body.(*dtos.AuthenticatedUserDto)
		if ok {
			ctx.Set("auth", authenticatedUserDto)
		}
	}
}
