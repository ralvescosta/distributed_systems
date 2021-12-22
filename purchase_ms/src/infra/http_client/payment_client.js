const { setTimeout } = require('timers/promises');

const { right } = require('../../domain/entities/either')

class PaymentClient {
  constructor(logger, telemetry) {
    this.logger = logger;
    this.telemetry = telemetry;
  }

  async payment({ context }) {
    const span = this.telemetry.createChildrenSpan({ context, name: 'HTTPS POST https://pagar.me.com.br/api/payments' })
    
    await setTimeout(200)
    
    span.end()
    return right(true)
  }
}

module.exports = { PaymentClient };