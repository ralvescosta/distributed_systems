const { ErrorCodeEnum } = require('../../domain/enums/error_code')

class ConflictError extends Error { 
  constructor(message, error) {
    super(message);
    this.name = "ConflictError";
    this.code = ErrorCodeEnum.ConflictError
    error ? this.stack = error.stack : undefined;
  }
}

module.exports = { ConflictError };