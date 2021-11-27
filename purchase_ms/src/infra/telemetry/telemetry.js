// const { NodeTracerProvider } = require('@opentelemetry/node')
const { trace } = require('@opentelemetry/api')
const { SimpleSpanProcessor, BasicTracerProvider } = require('@opentelemetry/tracing')
const { JaegerExporter } = require('@opentelemetry/exporter-jaeger')

class Telemetry {
  constructor(logger) {
    this.logger = logger;
  }

  start() {
    const tracer = new BasicTracerProvider()
  
    const exporter = new JaegerExporter({
      host: process.env.JAEGER_HOST,
      username: process.env.APP_NAME,
    })
    
    tracer.addSpanProcessor(new SimpleSpanProcessor(exporter))

    trace.setGlobalTracerProvider(tracer)
  }

  instrumentAmqp(queue, exchange, routingKey) {
    const cTracer = trace.getTracer(process.env.APP_NAME, '0.1.0');
    const span = cTracer.startSpan(`Queue: ${queue}`);

    span.setAttributes('Exchange', exchange)
    span.setAttributes('RoutingKey', routingKey)

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