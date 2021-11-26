const { right } = require('../../domain/entities/either')

class PurchaseUseCase {
  constructor(logger, inventoryClient, paymentClient, pubClient) {
    this.logger = logger;
    this.inventoryClient = inventoryClient;
    this.paymentClient = paymentClient;
    this.pubClient = pubClient;
  }

  async perform(order) {
    const isAvailable = await this.inventoryClient.verifyAvailability(order);
    if(isAvailable.isLeft()) {
      return isAvailable;
    }

    const orderAlreadyExist = await this.purchaseRepository.findByOrderId(order.orderId)
    if(orderAlreadyExist.isLeft()) {
      return orderAlreadyExist;
    }
    if(orderAlreadyExist.value) {
      return;
    }

    const payment = await this.paymentClient.payment(order);
    if(payment.isLeft()) {
      return payment;
    }

    const purchase = await this.purchaseRepository.create({...order, ...payment});
    if(purchase.isLeft()) {
      return purchase;
    }

    this.pubClient.updateInventory(order);
    this.pubClient.purchaseEmail(order);

    return right(true)
  }
}

module.exports = { PurchaseUseCase }