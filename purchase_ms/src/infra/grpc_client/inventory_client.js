class InventoryClient {
  constructor(logger) {
    this.logger = logger;
  }

  doSomething() {
    this.logger.info("[InventoryClient::doSomething]")
  }
}

module.exports = { InventoryClient };