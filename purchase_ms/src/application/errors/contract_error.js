const { ErrorCodeEnum } = require('../../domain/enums/error_code')

class ContractError extends Error { 
  constructor(message, error) {
    super(message);
    this.name = "ContractError";
    this.code = ErrorCodeEnum.ContractError
    error ? this.stack = error.stack : undefined;
  }
}

module.exports = { ContractError };