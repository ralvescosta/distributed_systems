const { Resource } = require('@opentelemetry/resources')
const { SemanticResourceAttributes } = require('@opentelemetry/semantic-conventions')
const { trace } = require('@opentelemetry/api')
const { BasicTracerProvider, BatchSpanProcessor } = require('@opentelemetry/tracing')
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger')

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

  instrumentAmqp(queue, exchange, routingKey) {
    const cTracer = trace.getTracer(process.env.APP_NAME, '0.1.0');
    const span = cTracer.startSpan(`Queue: ${queue}`);

    span.setAttribute("amqp.exchange", exchange)
    span.setAttribute("amqp.routingKey", routingKey)

    return span
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