const { Resource } = require('@opentelemetry/resources')
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions')
const { trace, context: apiContext, propagation } = require('@opentelemetry/api')
const { BasicTracerProvider, BatchSpanProcessor } = require('@opentelemetry/tracing')
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger')
const { Metadata } = require('@grpc/grpc-js');

class Telemetry {
  constructor(logger) {
    this.logger = logger;
    this.appName = process.env.APP_NAME || 'purchase-ms';
    this.appVersion = process.env.APP_VERSION || '0.0.1'
  }

  start() {
    const tracer = new BasicTracerProvider({
      resource: new Resource({
        [SemanticResourceAttributes.SERVICE_NAME]: this.appName
      })
    })
    const exporter = new JaegerExporter({
      host: process.env.JAEGER_HOST,
    })
    tracer.addSpanProcessor(new BatchSpanProcessor(exporter))
    trace.setGlobalTracerProvider(tracer)
  }

  instrumentAmqp({ queue, exchange, routingKey}) {
    const cTracer = trace.getTracer(this.appName, this.appVersion);
    const span = cTracer.startSpan(`Queue: ${queue}`);

    span.setAttribute("amqp.exchange", exchange)
    span.setAttribute("amqp.routingKey", routingKey)

    const context = trace.setSpan(apiContext.active(), span)
    return { span, context }
  }

  createChildrenSpan({ name, context }) {
    const cTracer = trace.getTracer(this.appName, this.appVersion);
    
    return cTracer.startSpan(name, undefined, context)
  }

  grpcInjector({ grpcCall = '', context }) {
    const cTracer = trace.getTracer(this.appName, this.appVersion);

    const span = cTracer.startSpan(grpcCall, undefined, context)
    const spanContext = span.spanContext()
    
    const metadata = new Metadata()
    metadata.add("traceparent", `00-${spanContext.traceId}-${spanContext.spanId}-01`)

    return { span, metadata }
  }

  basicExtractor() {
    return {
      keys: (carrier) =>  {
        return Object.keys(carrier).map((key) => [key, carrier[key]])
      },
      get: (carrier, key) => {
        return carrier[key]
      }
    }
  }

  amqpExtractor({ headers = {}, queue, exchange, routingKey }) {
    const { traceparent } = headers
    if(!traceparent)
      return this.instrumentAmqp({ queue, exchange, routingKey })

    const cTracer = trace.getTracer(this.appName, this.appVersion);

    const span = cTracer.startSpan(`amqpSubQueue: ${queue}`)
    let context = trace.setSpan(apiContext.active(), span)
    context = propagation.extract(
      context,
      headers,
      this.basicExtractor
    );
    
    span.setAttribute("amqp.exchange", exchange)
    span.setAttribute("amqp.routingKey", routingKey)

    return { span, context }
  }

  amqpInjector({ queue, exchange, routingKey, context }) {
    const cTracer = trace.getTracer(this.appName, this.appVersion);
    const span = cTracer.startSpan(`amqpPubQueue: ${queue}`, undefined, context)

    span.setAttribute("amqp.exchange", exchange)
    span.setAttribute("amqp.routingKey", routingKey)

    const spanContext = span.spanContext()

    return { span,  headers: { traceparent: `00-${spanContext.traceId}-${spanContext.spanId}-01`} }
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