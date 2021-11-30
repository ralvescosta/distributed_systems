const { Resource } = require('@opentelemetry/resources')
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions')
const { trace, context: apiContext } = require('@opentelemetry/api')
const { BasicTracerProvider, BatchSpanProcessor } = require('@opentelemetry/tracing')
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger')
const { Metadata } = require('@grpc/grpc-js');

class Telemetry {
  constructor(logger) {
    this.logger = logger;
  }

  start() {
    const tracer = new BasicTracerProvider({
      resource: new Resource({
        [SemanticResourceAttributes.SERVICE_NAME]: process.env.APP_NAME
      })
    })
    const exporter = new JaegerExporter({
      host: process.env.JAEGER_HOST,
    })
    tracer.addSpanProcessor(new BatchSpanProcessor(exporter))
    trace.setGlobalTracerProvider(tracer)
  }

  instrumentAmqp({ queue, exchange, routingKey}) {
    const cTracer = trace.getTracer(process.env.APP_NAME, process.env.APP_VERSION);
    const span = cTracer.startSpan(`Queue: ${queue}`);

    span.setAttribute("amqp.exchange", exchange)
    span.setAttribute("amqp.routingKey", routingKey)

    const context = trace.setSpan(apiContext.active(), span)
    return { span, context }
  }

  createChildrenSpan({ name, context }) {
    const cTracer = trace.getTracer(process.env.APP_NAME, process.env.APP_VERSION);
    
    return cTracer.startSpan(name, undefined, context)
  }

  grpcInjector({ grpcCall, context }) {
    const cTracer = trace.getTracer(process.env.APP_NAME, process.env.APP_VERSION);

    const span = cTracer.startSpan(grpcCall, undefined, context)
    const spanContext = span.spanContext()
    
    const metadata = new Metadata()
    metadata.add("traceparent", `00-${spanContext.traceId}-${spanContext.spanId}-01`)

    return { span, metadata }
  }

  amqpExtractor({ headers = {}, queue, exchange, routingKey }) {
    const { traceparent } = headers
    if(!traceparent)
      return this.instrumentAmqp({ queue, exchange, routingKey })

    const [, traceId, spanId] = traceparent.split('-')

    const span = trace.wrapSpanContext({
      spanId,
      traceId,
    })

    span.updateName(queue)
    span.setAttribute("amqp.exchange", exchange)
    span.setAttribute("amqp.routingKey", routingKey)

    const context = trace.setSpan(apiContext.active(), span)

    return { span, context }
  }

  getTracer() {
    return this.tracer
  }

  getSpan() {
    return this.tracer.getCurrentSpan()
  }

  getContext() {
    return this.tracer.getCurrentSpan().context()
  }
}

module.exports = { Telemetry }