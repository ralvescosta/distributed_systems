package interfaces

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type ITelemetry interface {
	GinMiddle() gin.HandlerFunc
	InstrumentQuery(ctx context.Context, sqlType string, sql string) opentracing.Span
	StartSpanFromRequest(header http.Header) opentracing.Span
	Inject(span opentracing.Span, request *http.Request) error
	Extract(header http.Header) (opentracing.SpanContext, error)
	Dispatch()
}
