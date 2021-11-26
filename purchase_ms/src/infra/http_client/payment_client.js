const { right } = require('../../domain/entities')

class PaymentClient {
  constructor(logger) {
    this.logger = logger;
  }

  payment() {
    this.logger.info("[PaymentClient::payment]")

    return right(true)
  }
}

module.exports = { PaymentClient };