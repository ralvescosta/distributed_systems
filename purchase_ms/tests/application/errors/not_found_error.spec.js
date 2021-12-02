const { NotFoundError } = require('../../../src/application/errors/not_found_error')
const { ErrorCodeEnum } = require('../../../src/domain/enums/error_code')

describe("Application :: Errors :: NotFoundError", () => {
  it("should create a new instance of NotFoundError", () => {
    const error = new NotFoundError()
    expect(error.name).toEqual(NotFoundError.name)
  })

  it("should contains the error code correctly", () => {
    const error = new NotFoundError()
    expect(error.code).toEqual(ErrorCodeEnum.NotFoundError)
  })
})