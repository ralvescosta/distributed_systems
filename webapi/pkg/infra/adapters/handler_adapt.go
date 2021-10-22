package adapters

import (
	"net/http"

	"webapi/pkg/app/interfaces"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

func HandlerAdapt(handler func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse, logger interfaces.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		request, err := GetHttpRequest(ctx)
		if err != nil {
			logger.Error("error while read request bytes")
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		result := handler(request)

		ctx.JSON(result.StatusCode, result.Body)
	}
}
