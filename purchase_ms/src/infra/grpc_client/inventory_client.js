const { promisify } = require('util');
const { right, left } = require('../../domain/entities/either')
const { GrpcClient } = require('./grpc_client')

class InventoryClient extends GrpcClient {
  constructor(logger) {
    super();
    this.logger = logger;
    this.loadProto('inventory.proto')
  }

  async verifyAvailability({ productId }) { //context
    const client = this.getClientInstance('inventory', 'Inventory', process.env.INVENTORY_MS_URL)
    
    if(client.isLeft())
      return client

    try {
      const result = await promisify(client.value.getProductById.bind(client.value))({ id: productId })
      return right(result)
    }catch(err) {
      return left(this.errorMapper(err))
    }
  }
}

module.exports = { InventoryClient };