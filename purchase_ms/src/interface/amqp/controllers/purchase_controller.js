class PurchaseController {
  constructor(logger) {
    this.logger = logger;
  }

  handle() {
    this.logger.info("[PurchaseController::handle]")
    return true
  }
}

module.exports = { PurchaseController };