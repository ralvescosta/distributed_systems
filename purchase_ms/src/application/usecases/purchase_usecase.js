class PurchaseUseCase {
  constructor(logger, inventoryClient, paymentClient) {
    this.logger = logger;
    this.inventoryClient = inventoryClient;
    this.paymentClient = paymentClient;
  }

  perform() {
    this.logger.info("[PurchaseUseCase::perform]");
    
    this.inventoryClient.doSomething();
    this.paymentClient.doSomething();
  }
}

module.exports = { PurchaseUseCase }