class PubClient {
  constructor(logger, messageBroker) {
    this.logger = logger;
    this.messageBroker = messageBroker;
  }

  updateInventory({ order, payment, context }) {}

  purchaseEmail({ order, payment, context }) {}
}

module.exports = { PubClient }