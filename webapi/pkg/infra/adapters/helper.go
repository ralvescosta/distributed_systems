package adapters

import (
	"bytes"
	"io/ioutil"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
)

func GetHttpRequest(ctx *gin.Context) (internalHttp.HttpRequest, error) {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return internalHttp.HttpRequest{}, err
	}
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	params := make(map[string]string)
	for _, param := range ctx.Params {
		params[param.Key] = param.Value
	}

	auth, _ := ctx.Get("auth")

	return internalHttp.HttpRequest{
		Body:    body,
		Headers: ctx.Request.Header,
		Params:  params,
		Auth:    auth,
		Ctx:     ctx.Request.Context(),
		Txn:     nrgin.Transaction(ctx),
	}, nil
}
