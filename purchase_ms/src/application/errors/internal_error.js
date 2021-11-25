const { ErrorCodeEnum } = require('../../domain/enums/error_code')

class InternalError extends Error { 
  constructor(message, error) {
    super(message);
    this.name = "InternalError";
    this.code = ErrorCodeEnum.InternalError
    error ? this.stack = error.stack : undefined;
  }
}

module.exports = { InternalError };