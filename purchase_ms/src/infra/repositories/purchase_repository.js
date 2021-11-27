const { InternalError } = require('../../application/errors/internal_error')
const { right, left } = require('../../domain/entities/either')
const { PurchaseModel } = require('../database/purchase_model')

class PurchaseRepository {
  constructor(logger, dbConnection) {
    this.logger = logger;
    this.dbConnection = dbConnection;
  }

  findByOrderId({ orderId, context }) {
    // try {
    //   const result = PurchaseModel.findOne({ orderId })
    //   return right(result)
    // }catch (err) {
    //   return left(new InternalError("Error while find purchase by orderId", err))
    // }
  }

  create({ order, payment, context }) {
    // try {
    //   const result = PurchaseModel.create(purchase)
    //   return right(result)
    // }catch (err) {
    //   return left(new InternalError("Error while try to create purchase", err))
    // }
  }
}

module.exports = { PurchaseRepository }