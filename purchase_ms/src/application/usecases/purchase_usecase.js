const { right } = require('../../domain/entities/either')

class PurchaseUseCase {
  constructor(logger, inventoryClient, paymentClient, pubClient) {
    this.logger = logger;
    this.inventoryClient = inventoryClient;
    this.paymentClient = paymentClient;
    this.pubClient = pubClient;
  }

  async perform(order, context) {
    const isAvailable = await this.inventoryClient.verifyAvailability({ order, context });
    if(isAvailable.isLeft()) {
      return isAvailable;
    }

    const orderAlreadyExist = await this.purchaseRepository.findByOrderId({ orderId: order.orderId, context })
    if(orderAlreadyExist.isLeft()) {
      return orderAlreadyExist;
    }
    if(orderAlreadyExist.value) {
      return;
    }

    const payment = await this.paymentClient.payment({ order, context });
    if(payment.isLeft()) {
      return payment;
    }

    const purchase = await this.purchaseRepository.create({ order, payment, context });
    if(purchase.isLeft()) {
      return purchase;
    }

    this.pubClient.updateInventory({ order, payment, context });
    this.pubClient.purchaseEmail({ order, payment, context });

    return right(true)
  }
}

module.exports = { PurchaseUseCase }