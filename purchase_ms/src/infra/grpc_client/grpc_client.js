const { loadPackageDefinition, credentials, status } = require('@grpc/grpc-js');

const protoLoader = require('@grpc/proto-loader');
const { resolve } = require('path')

const { left, right } = require('../../domain/entities/either')
const { UnavailableServiceError } = require('../../application/errors/unavailable_service_error')
const { InternalError } = require('../../application/errors/internal_error')
const { NotFoundError } = require('../../application/errors/not_found_error')

class GrpcClient {
  constructor() {
    this.packageDefinition
  }

  loadProto(protoFile) {
    this.packageDefinition = protoLoader.loadSync(
      resolve(__dirname, 'proto', protoFile),
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
    const definition = loadPackageDefinition(this.packageDefinition)
    
    const grpcPackage = definition[pkg]
    if (!grpcPackage)
      return left(new UnavailableServiceError("Inventory Client - Unavailable service"))

    const GrpcClient = grpcPackage[client]
    if (!client) 
      return left(new UnavailableServiceError("Inventory Client - Unavailable service"))

    try {
      return right(new GrpcClient(host, credentials.createInsecure()))
    } catch (err) {
      return left(new UnavailableServiceError("Inventory Client - Unavailable service"))
    }
  }

  errorMapper(err) {
    switch (err.code) {
      case status.NOT_FOUND:
        return left(new NotFoundError("Register not found", err))
      default:
        return left(new InternalError("Some error occur", err))
    }
  }
}

module.exports = { GrpcClient }