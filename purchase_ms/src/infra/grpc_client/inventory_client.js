const { promisify } = require('util');
const { right } = require('../../domain/entities/either')
const { GrpcClient } = require('./grpc_client')

class InventoryClient extends GrpcClient {
  constructor(logger) {
    super();
    this.logger = logger;
    this.loadProto('inventory.proto')
  }

  async verifyAvailability(id) {
    const client = this.getClientInstance('inventory', 'Inventory', process.env.INVENTORY_MS_URL)
    
    if(client.isLeft())
      return client

    const result = await promisify(client.value.getProductById)({ id })
    this.logger.info("[InventoryClient::verifyAvailability]", result)
    return right(true)
  }
}

module.exports = { InventoryClient };