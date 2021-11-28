const { ErrorCodeEnum } = require('../../domain/enums/error_code')

class NotFoundError extends Error { 
  constructor(message, error) {
    super(message);
    this.name = "NotFoundError";
    this.code = ErrorCodeEnum.NotFoundError
    error ? this.stack = error.stack : undefined;
  }
}

module.exports = { NotFoundError };