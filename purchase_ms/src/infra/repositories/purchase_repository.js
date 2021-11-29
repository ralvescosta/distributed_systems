// const { InternalError } = require('../../application/errors/internal_error')
const { right, left } = require('../../domain/entities/either')
// const { PurchaseModel } = require('../database/purchase_model')

class PurchaseRepository {
  constructor(logger, dbConnection, telemetry) {
    this.logger = logger;
    this.dbConnection = dbConnection;
    this.telemetry = telemetry;
  }

  findByOrderId({ context }) {
    const span = this.telemetry.createChildrenSpan({ context, name: 'MONGO GET BY ORDER_ID' })
    // try {
    //   const result = PurchaseModel.findOne({ orderId })
    //   return right(result)
    // }catch (err) {
    //   return left(new InternalError("Error while find purchase by orderId", err))
    // }
    span.end()
    return right(undefined)
  }

  create({ context }) { //order, payment
    const span = this.telemetry.createChildrenSpan({ context, name: 'MONGO CREATE PURCHASE' })
    // try {
    //   const result = PurchaseModel.create(purchase)
    //   return right(result)
    // }catch (err) {
    //   return left(new InternalError("Error while try to create purchase", err))
    // }
    span.end()
    return right({})
  }
}

module.exports = { PurchaseRepository }