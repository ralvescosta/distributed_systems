const { InternalError } = require('../../application/errors/internal_error')
const { right, left } = require('../../domain/entities/either')
class PurchaseRepository {
  constructor(logger, dbConnection) {
    this.logger = logger;
    this.dbConnection = dbConnection;
  }

  findByOrderId(orderId) {
    try {
      const result = this.dbConnection.findOne({ orderId})
      return right(result)
    }catch (err) {
      return left(new InternalError("Error while find purchase by orderId", err))
    }
  }

  create(purchase) {
    try {
      const result = this.dbConnection.create(purchase)
      return right(result)
    }catch (err) {
      return left(new InternalError("Error while try to create purchase", err))
    }
  }
}

module.exports = { PurchaseRepository }