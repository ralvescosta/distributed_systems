const { PurchaseUseCase } = require('../../../src/application/usecases/purchase_usecase')
const { right } = require('../../../src/domain/entities/either')

const makeSut = () => {
  const loggerSpy = {}
  const inventoryClientSpy = {
    verifyAvailability: jest.fn(() => right(true)),
  }
  const purchaseRepository = {
    findByOrderId: jest.fn(() => right(false)),
    create: jest.fn(() => right(true)),
  }
  const paymentClient = {
    payment: jest.fn(() => right(true)),
  }
  const pubClient = {
    updateInventory: jest.fn(() => right(true)),
    purchaseEmail: jest.fn(() => right(true)),
  }

  const sut = new PurchaseUseCase(loggerSpy, inventoryClientSpy, purchaseRepository, paymentClient, pubClient)
  
  return {
    sut,
    loggerSpy,
    inventoryClientSpy,
    purchaseRepository,
    paymentClient,
    pubClient,
  }
}

describe('APPLICATION :: USECASES :: PurchaseUseCase', () => {
  describe('perform()', () => {
    it('should execute the purchase correctly', async () => {
      const { sut, inventoryClientSpy, purchaseRepository, paymentClient, pubClient } = makeSut()
      const order = {}, context = {}

      const result = await sut.perform({ order, context })

      expect(result.isRight()).toBeTruthy()
      expect(inventoryClientSpy.verifyAvailability).toHaveBeenCalledTimes(1)
      expect(purchaseRepository.findByOrderId).toHaveBeenCalledTimes(1)
      expect(paymentClient.payment).toHaveBeenCalledTimes(1)
      expect(purchaseRepository.create).toHaveBeenCalledTimes(1)
      expect(pubClient.updateInventory).toHaveBeenCalledTimes(1)
      expect(pubClient.purchaseEmail).toHaveBeenCalledTimes(1)
    })
  })
})