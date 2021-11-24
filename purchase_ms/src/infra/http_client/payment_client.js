class PaymentClient {
  constructor(logger) {
    this.logger = logger;
  }

  doSomething() {
    this.logger.info("[PaymentClient::doSomething]")
  }
}

module.exports = { PaymentClient };