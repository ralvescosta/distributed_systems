const { left } = require('../../../domain/entities/either')
class PurchaseController {
  constructor(logger, purchaseUseCase, telemetry) {
    this.logger = logger;
    this.purchaseUseCase = purchaseUseCase;
    this.telemetry = telemetry;
  }

  async handle({ content }) { //fields, properties
    const span = this.telemetry.instrumentAmqp(
      process.env.AMQP_QUEUE, 
      process.env.AMQP_EXCHANGE, 
      process.env.AMQP_ROUTING_KEY,
    )

    let order = {}
    try {
      order = JSON.parse(content)
    }catch(err) {
      this.logger.error("[PurchaseController::handle]")
    }
    
    const result = await this.purchaseUseCase.perform({ order, context: span.spanContext() })
    if (result.isLeft()) {
      span.setAttribute("error", true)
    }

    span.end();
    return left({ error_code: 40 })
  }
}

module.exports = { PurchaseController };