const { PurchaseModel } = require('../../../src/infra/database/purchase_model')

describe('INFRA :: DATABASE :: PurchaseModel', () => {
  beforeEach(() => {
    jest.clearAllMocks()
  })
  it('Should PurchaseModel exist', () => {
    expect(PurchaseModel.modelName).toBe('Purchase')
  })
})