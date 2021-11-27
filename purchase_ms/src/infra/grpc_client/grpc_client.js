const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { resolve } = require('path')

const { left, right } = require('../../domain/entities/either')
const { UnavailableServiceError } = require('../../application/errors/unavailable_service_error')

class GrpcClient {
  constructor() {
    this.packageDefinition
  }

  loadProto(protoFile) {
    this.packageDefinition = protoLoader.loadSync(
      resolve(__dirname, '..', 'proto', protoFile),
      {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
      },
    )
  }

  getClientInstance(pkg, client, host) {
    const definition = grpc.loadPackageDefinition(this.packageDefinition)
    
    const grpcPackage = definition[pkg]
    if (!grpcPackage)
      return left(new UnavailableServiceError("Inventory Client - Unavailable service"))

    const ClientClass = grpcPackage[client]
    if (!client) 
      return left(new UnavailableServiceError("Inventory Client - Unavailable service"))

    try {
      right(new ClientClass(host, grpc.credentials.createInsecure()))
    } catch (err) {
      return left(new UnavailableServiceError("Inventory Client - Unavailable service"))
    }
    
  }
}

module.exports = { GrpcClient }