class PurchaseSubscriber {
  constructor(logger, messageBroker, purchaseController){
    this.logger = logger;
    this.messageBroker = messageBroker;
    this.purchaseController = purchaseController;
  }

  subscribe() {
    this.logger.info("[PurchaseSubscriber::subscribe]")
    const AMQP_QUEUE = process.env.AMQP_QUEUE

    this.messageBroker.sub(
      AMQP_QUEUE,
      this.purchaseController,
      { noAck: false },
    )
  }
}

module.exports = { PurchaseSubscriber };