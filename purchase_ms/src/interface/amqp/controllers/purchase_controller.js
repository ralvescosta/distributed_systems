const { left } = require('../../../domain/entities/either')
class PurchaseController {
  constructor(logger, purchaseUseCase, telemetry) {
    this.logger = logger;
    this.purchaseUseCase = purchaseUseCase;
    this.telemetry = telemetry;
  }

  handle({ content }) { //fields, properties
    const span = this.telemetry.instrumentAmqp(
      process.env.AMQP_QUEUE, 
      process.env.AMQP_EXCHANGE, 
      process.env.AMQP_ROUTING_KEY,
    )

    try {
      const order = JSON.parse(content)
      this.purchaseUseCase.perform({ order, context: span.spanContext() })
    }catch(err) {
      this.logger.error("[PurchaseController::handle]")
    }
    
    this.logger.info("[PurchaseController::handle]")

    span.end();
    return left({ error_code: 40 })
  }
}

module.exports = { PurchaseController };