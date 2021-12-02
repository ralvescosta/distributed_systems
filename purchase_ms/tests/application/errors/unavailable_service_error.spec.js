const { UnavailableServiceError } = require('../../../src/application/errors/unavailable_service_error')
const { ErrorCodeEnum } = require('../../../src/domain/enums/error_code')

describe("Application :: Errors :: UnavailableServiceError", () => {
  it("should create a new instance of UnavailableServiceError", () => {
    const error = new UnavailableServiceError()
    expect(error.name).toEqual(UnavailableServiceError.name)
  })

  it("should contains the error code correctly", () => {
    const error = new UnavailableServiceError()
    expect(error.code).toEqual(ErrorCodeEnum.UnavailableServiceError)
  })
})