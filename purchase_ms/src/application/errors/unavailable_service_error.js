const { ErrorCodeEnum } = require('../../domain/enums/error_code')

class UnavailableServiceError extends Error { 
  constructor(message, error) {
    super(message);
    this.name = "UnavailableServiceError";
    this.code = ErrorCodeEnum.UnavailableServiceError
    error ? this.stack = error.stack : undefined;
  }
}

module.exports = { UnavailableServiceError };