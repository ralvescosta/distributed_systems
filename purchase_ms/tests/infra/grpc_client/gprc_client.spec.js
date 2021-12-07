const { GrpcClient } = require('../../../src/infra/grpc_client/grpc_client')
const { NotFoundError } = require('../../../src/application/errors/not_found_error')
const protoLoader = require('@grpc/proto-loader')
const path = require('path')
const { loadPackageDefinition, status } = require('@grpc/grpc-js');

jest.mock('@grpc/proto-loader', () => ({
  loadSync: jest.fn()
}))
jest.mock('path', () => ({
  resolve: jest.fn()
}))
jest.mock('@grpc/grpc-js', () => ({
  loadPackageDefinition: jest.fn(() => ({
    'some-package': {
      'some-client': class SomeClass {}
    }
  })),
  credentials: {
    createInsecure: jest.fn
  },
  status: {
    NOT_FUND: 40
  }
}))
describe('INFRA :: GRPC CLIENT :: GrpcClient', () => {
  describe('loadProto()', () => {
    beforeEach(() => {
      jest.clearAllMocks()
    })
    it('should load proto correctly', () => {
      const sut = new GrpcClient()

      sut.loadProto('some_file')

      expect(protoLoader.loadSync).toHaveBeenCalledTimes(1)
      expect(path.resolve).toHaveBeenCalledTimes(1)
    })
  })

  describe('getClientInstance()', () => {
    beforeEach(() => {
      jest.clearAllMocks()
    })
    it('should return the client instance correctly', () => {
      const sut = new GrpcClient()
    
      const pkg = 'some-package'
      const client = 'some-client'
      const host = 'some-host'

      const result = sut.getClientInstance(pkg, client, host)

      expect(result.isRight()).toBeTruthy()
      expect(loadPackageDefinition).toHaveBeenCalledTimes(1)
    })
  })

  describe('errorMapper()', () => {
    beforeEach(() => {
      jest.clearAllMocks()
    })
    it('should map grpc error', () => {
      const sut = new GrpcClient()

      const result = sut.errorMapper({ code: status.NOT_FOUND })

      expect(result.isLeft()).toBeTruthy()
      expect(result.value).toBeInstanceOf(NotFoundError)
    })
  })
})