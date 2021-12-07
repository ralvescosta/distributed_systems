const { InventoryClient } = require('../../../src/infra/grpc_client/inventory_client')
const { right } = require('../../../src/domain/entities/either')

const makeSut = () => {
  const loggerSpy = {}
  const telemetrySpy = {
    grpcInjector: jest.fn(() => ({
      span: {
        end: jest.fn(),
        setAttribute: jest.fn(),
        metadata: {}
      }
    }))
  }

  const clientInstance = new class ClientInstance{
    getProductById(data, metadata, cb) {
      return cb()
    }
  }

  const sut = new InventoryClient(loggerSpy, telemetrySpy)

  jest.spyOn(sut, sut.loadProto.name)
  jest.spyOn(sut, sut.getClientInstance.name).mockReturnValueOnce(right(clientInstance))
  jest.spyOn(sut, sut.errorMapper.name).mockReturnValueOnce(right(clientInstance))

  return {
    sut,
    clientInstance,
  }
}

describe('INFRA :: GRPC CLIENT :: InventoryClient', () => {
  describe('verifyAvailability()', () => {
    beforeEach(() => {
      jest.clearAllMocks()
    })
    it('should execute correctly', async () => {
      const { sut } = makeSut()

      const productId = 'some-id'
      const context = {}

      const result = await sut.verifyAvailability({ productId, context })

      expect(result.isRight()).toBeTruthy()
      expect(sut.getClientInstance).toHaveBeenCalledTimes(1)
    })
  })
})