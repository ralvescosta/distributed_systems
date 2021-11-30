const { left } = require('../../../domain/entities/either')
class PurchaseController {
  constructor(logger, purchaseUseCase, telemetry) {
    this.logger = logger;
    this.purchaseUseCase = purchaseUseCase;
    this.telemetry = telemetry;
  }

  async handle({ content, properties }) { //fields, properties
    const { span, context } = this.telemetry.amqpExtractor({
      headers: properties.headers,
      queue: process.env.AMQP_QUEUE, 
      exchange: process.env.AMQP_EXCHANGE, 
      routingKey: process.env.AMQP_ROUTING_KEY,
    })

    let order = {}
    try {
      order = JSON.parse(content)
    }catch(err) {
      this.logger.error("[PurchaseController::handle]")
    }
    
    const result = await this.purchaseUseCase.perform({ order, context })
    if (result.isLeft()) {
      span.setAttribute("error", true)
    }

    span.end();
    return left({ error_code: 40 })
  }
}

module.exports = { PurchaseController };