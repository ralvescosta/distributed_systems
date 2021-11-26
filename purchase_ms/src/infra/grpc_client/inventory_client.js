const { right } = require('../../domain/entities/either')

class InventoryClient {
  constructor(logger) {
    this.logger = logger;
  }

  verifyAvailability() {
    this.logger.info("[InventoryClient::verifyAvailability]")
    return right(true)
  }
}

module.exports = { InventoryClient };