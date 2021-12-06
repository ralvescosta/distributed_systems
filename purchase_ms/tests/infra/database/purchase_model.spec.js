const { PurchaseModel } = require('../../../src/infra/database/purchase_model')

describe('INFRA :: DATABASE :: PurchaseModel', () => {
  it('Should PurchaseModel exist', () => {
    expect(PurchaseModel.modelName).toBe('Purchase')
  })
})