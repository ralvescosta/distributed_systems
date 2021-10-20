package adapters

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"webapi/pkg/app/interfaces"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

var readAllBody = ioutil.ReadAll

func HandlerAdapt(handler func(httpRequest internalHttp.HttpRequest) internalHttp.HttpResponse, logger interfaces.ILogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		body, err := readAllBody(ctx.Request.Body)
		if err != nil {
			logger.Error("error while read request bytes")
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		params := make(map[string]string)
		for _, param := range ctx.Params {
			params[param.Key] = param.Value
		}

		request := internalHttp.HttpRequest{
			Body:    body,
			Headers: ctx.Request.Header,
			Params:  params,
			Ctx:     ctx.Request.Context(),
			Txn:     nrgin.Transaction(ctx),
		}

		result := handler(request)

		ctx.JSON(result.StatusCode, result.Body)
	}
}
