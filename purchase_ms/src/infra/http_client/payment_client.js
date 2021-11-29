const { right } = require('../../domain/entities/either')

class PaymentClient {
  constructor(logger, telemetry) {
    this.logger = logger;
    this.telemetry = telemetry;
  }

  payment({ context }) {
    const span = this.telemetry.createChildrenSpan({ context, name: 'HTTPS POST https://getnet.com.br/api/payments' })
    span.end()
    return right(true)
  }
}

module.exports = { PaymentClient };