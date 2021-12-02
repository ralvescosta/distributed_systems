const { InternalError } = require('../../../src/application/errors/internal_error')
const { ErrorCodeEnum } = require('../../../src/domain/enums/error_code')

describe("Application :: Errors :: InternalError", () => {
  it("should create a new instance of InternalError", () => {
    const error = new InternalError()
    expect(error.name).toEqual(InternalError.name)
  })

  it("should contains the error code correctly", () => {
    const error = new InternalError()
    expect(error.code).toEqual(ErrorCodeEnum.InternalError)
  })
})