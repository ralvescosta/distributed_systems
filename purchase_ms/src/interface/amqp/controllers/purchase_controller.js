const { left } = require('../../../domain/entities/either')
class PurchaseController {
  constructor(logger, purchaseUseCase) {
    this.logger = logger;
    this.purchaseUseCase = purchaseUseCase;
  }

  handle() {
    this.logger.info("[PurchaseController::handle]")
    this.purchaseUseCase.perform()
    return left({ error_code: 40 })
  }
}

module.exports = { PurchaseController };