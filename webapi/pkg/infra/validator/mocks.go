package validator

import "webapi/pkg/app/interfaces"

type validatorToTest struct {
	validate interfaces.IValidator
}

func newValidatorToTest() validatorToTest {
	validate := NewValidator()

	return validatorToTest{validate}
}
