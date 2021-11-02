package validator

import "webapi/pkg/app/interfaces"

type ValidatorToTest struct {
	validate interfaces.IValidator
}

func NewValidatorToTest() ValidatorToTest {
	validate := NewValidator()

	return ValidatorToTest{validate}
}
