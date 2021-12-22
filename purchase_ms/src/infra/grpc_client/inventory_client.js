const { promisify } = require('util');

const { right, left } = require('../../domain/entities/either')
const { GrpcClient } = require('./grpc_client')

class InventoryClient extends GrpcClient {
  constructor(logger, telemetry) {
    super();
    this.logger = logger;
    this.telemetry = telemetry;
    this.loadProto('inventory.proto')
  }

  async verifyAvailability({ productId, context }) {
    const { span, metadata } = this.telemetry.grpcInjector({ grpcCall: "grpcGetProductById", context })
    
    const client = this.getClientInstance('inventory', 'Inventory', process.env.INVENTORY_MS_URL)
    
    if(client.isLeft()){
      this.telemetry.handleError(span, client.value)
      return client
    }
      
    try {
      const result = await promisify(client.value.getProductById.bind(client.value))({ id: productId }, metadata)
      return right(result)
    } catch(err) {
      this.telemetry.handleError(span, err)
      return left(this.errorMapper(err))
    } finally {
      span.end()
    }
  }
}

module.exports = { InventoryClient };