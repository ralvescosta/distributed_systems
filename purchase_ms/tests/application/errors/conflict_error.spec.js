const { ConflictError } = require('../../../src/application/errors/conflict_error')
const { ErrorCodeEnum } = require('../../../src/domain/enums/error_code')

describe("Application :: Errors :: ConflictError", () => {
  it("should create a new instance of ConflictError", () => {
    const error = new ConflictError()
    expect(error.name).toEqual(ConflictError.name)
  })

  it("should contains the error code correctly", () => {
    const error = new ConflictError()
    expect(error.code).toEqual(ErrorCodeEnum.ConflictError)
  })
})