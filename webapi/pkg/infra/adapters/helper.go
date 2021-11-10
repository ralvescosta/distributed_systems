package adapters

import (
	"bytes"
	"context"
	"io/ioutil"
	internalHttp "webapi/pkg/interfaces/http"

	"github.com/gin-gonic/gin"
)

func GetHttpRequest(ginCtx *gin.Context) (internalHttp.HttpRequest, error) {
	body, err := ioutil.ReadAll(ginCtx.Request.Body)
	if err != nil {
		return internalHttp.HttpRequest{}, err
	}
	ginCtx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	params := make(map[string]string)
	for _, param := range ginCtx.Params {
		params[param.Key] = param.Value
	}

	auth, _ := ginCtx.Get("auth")
	tracerCtx, _ := ginCtx.Get("tracerCtx")

	return internalHttp.HttpRequest{
		Body:    body,
		Headers: ginCtx.Request.Header,
		Params:  params,
		Auth:    auth,
		Ctx:     tracerCtx.(context.Context),
	}, nil
}
