package adapters

import (
	"net/http"

	"webapi/pkg/app/interfaces"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

func HandlerAdapt(handler func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse, logger interfaces.ILogger) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		request, err := GetHttpRequest(ginCtx)
		if err != nil {
			logger.Error("error while read request bytes")
			ginCtx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
			return
		}

		result := handler(request)

		ginCtx.JSON(result.StatusCode, result.Body)
	}
}
