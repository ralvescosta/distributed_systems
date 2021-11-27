const { right } = require('../../domain/entities/either')

class PaymentClient {
  constructor(logger) {
    this.logger = logger;
  }

  payment({ orderId, context }) {
    return right(true)
  }
}

module.exports = { PaymentClient };