package telemetry

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"webapi/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type telemetry struct {
	tracer opentracing.Tracer
	closer io.Closer
}

func NewTelemetry() interfaces.ITelemetry {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: failed to read config from env vars: %v\n", err))
	}
	//config.Reporter(jaeger.NewLoggingReporter(jaeger.StdLogger))
	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)

	return &telemetry{tracer, closer}
}

func (pst *telemetry) Dispatch() {
	pst.closer.Close()
}

func (telemetry) GinMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
	}
}

const (
	TAG_SQL_SELECT_USER = "SQL SELECT USER"
	TAG_SQL_INSERT_USER = "SQL INSERT USER"
)

func (pst *telemetry) InstrumentQuery(ctx context.Context, sqlType string, sql string) opentracing.Span {
	span := opentracing.SpanFromContext(ctx)
	span = pst.tracer.StartSpan(sqlType, opentracing.ChildOf(span.Context()))
	ext.SpanKindRPCClient.Set(span)
	ext.PeerService.Set(span, "postgreSQL")
	span.SetTag("sql.query", sql)

	return span
}

func (pst *telemetry) InstrumentGRPCClient(ctx context.Context, clientName string) (opentracing.Span, context.Context) {
	span := opentracing.SpanFromContext(ctx)
	span = pst.tracer.StartSpan(clientName, opentracing.ChildOf(span.Context()))
	ext.SpanKindRPCClient.Set(span)
	ext.PeerService.Set(span, "gRPC Client")
	return span, opentracing.ContextWithSpan(ctx, span)
}

// StartSpanFromRequest extracts the parent span context from the inbound HTTP request
// and starts a new child span if there is a parent span.
func (pst *telemetry) StartSpanFromRequest(header http.Header) opentracing.Span {
	spanCtx, _ := pst.Extract(header)
	return pst.tracer.StartSpan("ping-receive", ext.RPCServerOption(spanCtx))
}

// Inject injects the outbound HTTP request with the given span's context to ensure
// correct propagation of span context throughout the trace.
func (telemetry) Inject(span opentracing.Span, request *http.Request) error {
	return span.Tracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(request.Header))
}

// Extract extracts the inbound HTTP request to obtain the parent span's context to ensure
// correct propagation of span context throughout the trace.
func (pst *telemetry) Extract(header http.Header) (opentracing.SpanContext, error) {
	return pst.tracer.Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(header))
}

func (pst *telemetry) GetTracer() opentracing.Tracer {
	return pst.tracer
}
