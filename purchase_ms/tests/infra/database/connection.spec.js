const { DbConnection } = require('../../../src/infra/database/connection')
const { InternalError } = require('../../../src/application/errors/internal_error')
const mongoose = require('mongoose');

const makeSut = () => {
  process.env.MONGO_CONNECTION_URI = 'some_uri'

  const loggerSpy = {
    info: jest.fn(),
    error: jest.fn()
  }
  const sut = new DbConnection(loggerSpy)

  return { sut, loggerSpy }
}

jest.mock('mongoose', () => ({
  connect: jest.fn((uri, options, cb) => {
    if(uri === undefined) {
      return cb(new Error()) 
    }
    cb()
  })
}))

describe('INFRA :: DATABASE :: Connection', () => {
  describe('INFRA :: DATABASE :: Connection :: connect', () => {
    it('should connect in mongo db successfully', async () => {
      const { sut, loggerSpy } = makeSut()

      const result = await sut.connect();

      expect(result.isRight()).toBeTruthy()
      expect(mongoose.connect).toHaveBeenCalledTimes(1)
      expect(loggerSpy.info).toHaveBeenCalledTimes(1)
    })

    it('should return Left with InternalError if some error occur in mongo connection', async () => {
      const { sut, loggerSpy } = makeSut()
      delete process.env.MONGO_CONNECTION_URI

      const result = await sut.connect();

      expect(result.isLeft()).toBeTruthy()
      expect(mongoose.connect).toHaveBeenCalledTimes(1)
      expect(loggerSpy.error).toHaveBeenCalledTimes(1)
      expect(result.value).toBeInstanceOf(InternalError)
    })
  })
})