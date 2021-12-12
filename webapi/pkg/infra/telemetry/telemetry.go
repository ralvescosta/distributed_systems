package telemetry

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"google.golang.org/grpc/metadata"
)

type ITelemetry interface {
	GinMiddle() gin.HandlerFunc
	InstrumentQuery(ctx context.Context, sqlType string, sql string) opentracing.Span
	InstrumentGRPCClient(ctx context.Context, clientName string) (opentracing.Span, context.Context)
	InstrumentAMQPPublisher(ctx context.Context, exchangeName, queueName string) (opentracing.Span, context.Context)
	StartSpanFromRequest(header http.Header) opentracing.Span
	Inject(span opentracing.Span, request *http.Request) error
	Extract(header http.Header) (opentracing.SpanContext, error)
	Dispatch()
	GetTracer() opentracing.Tracer
}

type telemetry struct {
	tracer opentracing.Tracer
	closer io.Closer
}

func NewTelemetry() ITelemetry {
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

	jaegerCtx, _ := span.Context().(jaeger.SpanContext)

	ctxWithHeaders := metadata.NewOutgoingContext(
		opentracing.ContextWithSpan(ctx, span),
		metadata.Pairs("traceparent", fmt.Sprintf("00-%s-%s-01", jaegerCtx.ParentID(), jaegerCtx.SpanID())),
	)

	return span, ctxWithHeaders
}

func (pst *telemetry) InstrumentAMQPPublisher(ctx context.Context, exchangeName, queueName string) (opentracing.Span, context.Context) {
	span := opentracing.SpanFromContext(ctx)
	span = pst.tracer.StartSpan(fmt.Sprintf("exchange: %s", exchangeName), opentracing.ChildOf(span.Context()))

	ext.SpanKindRPCClient.Set(span)
	ext.PeerService.Set(span, "AMQP Pub")
	span.SetTag("amqp.exchange", exchangeName)
	span.SetTag("amqp.queue", queueName)

	jaegerCtx, _ := span.Context().(jaeger.SpanContext)

	ctxWithHeaders := metadata.NewOutgoingContext(
		opentracing.ContextWithSpan(ctx, span),
		metadata.Pairs("traceparent", fmt.Sprintf("00-%s-%s-01", jaegerCtx.ParentID(), jaegerCtx.SpanID())),
	)

	return span, ctxWithHeaders
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
