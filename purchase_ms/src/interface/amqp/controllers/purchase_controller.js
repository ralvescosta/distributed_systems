const { left, right } = require('../../../domain/entities/either')
const { ContractError } = require('../../../application/errors/contract_error')

class PurchaseController {
  constructor(logger, purchaseUseCase, telemetry) {
    this.logger = logger;
    this.purchaseUseCase = purchaseUseCase;
    this.telemetry = telemetry;
  }

  async handle({ content, properties }) {
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
      this.telemetry.handleError(span, err)
      return left(new ContractError("Error whiling serialize content to json", err))
    }
    
    const result = await this.purchaseUseCase.perform({ order, context })
    if (result.isLeft()) {
      this.telemetry.handleError(span, result.value)
      return result
    }
    
    span.end()
    return right(true)
  }
}

module.exports = { PurchaseController };