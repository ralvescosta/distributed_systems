const { left, right } = require('../../domain/entities/either')
const { InternalError } = require('../../application/errors/internal_error')

const mongoose = require('mongoose');
const { promisify } = require('util')
class DbConnection {
  constructor(logger) {
    this.logger = logger;
  }

  async connect() {
    const MONGO_CONNECTION_URI = process.env.MONGO_CONNECTION_URI

    try {
      await promisify(mongoose.connect)(
        MONGO_CONNECTION_URI, 
        { 
          useUnifiedTopology: true,
        })

        this.logger.info('Connected in Database!')
        return right(true)
    } catch (err) {
      this.logger.error(err)
      return left(new InternalError("Error whiling connect in MongoDb", err))
    }
  }
}

module.exports = { DbConnection }