class PubClient {
  constructor(logger, messageBroker, telemetry) {
    this.logger = logger;
    this.messageBroker = messageBroker;
    this.telemetry = telemetry;
  }

  updateInventory({ order, payment, context }) {
    const { span } = this.telemetry.amqpInjector({ 
      queue: 'updateInventoryQueue', 
      exchange: 'updateInventoryExchange', 
      routingKey: 'updateInventoryRoutingKey',
      context,
    })

    span.end()
  }

  purchaseEmail({ order, payment, context }) {
    const { span } = this.telemetry.instrumentAmqp({ 
      queue: 'purchaseEmailQueue', 
      exchange: 'purchaseEmailExchange', 
      routingKey: 'purchaseEmailRoutingKey',
      context,
    })

    span.end()
  }
}

module.exports = { PubClient }