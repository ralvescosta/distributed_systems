package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_ValidateStruct_Returns_Error_When_Some_Field_Is_Wrong(t *testing.T) {
	sut := newValidatorToTest()

	type SomeStruct struct {
		Email string `validate:"required,email"`
	}

	invalid := SomeStruct{
		Email: "email",
	}

	result := sut.validate.ValidateStruct(invalid)

	assert.NotNil(t, result)
}

func Test_Should_ValidateStruct_Returns_Nil_When_Some_Field_Is_Correctly(t *testing.T) {
	sut := newValidatorToTest()

	type SomeStruct struct {
		Email string `validate:"required,email"`
	}

	invalid := SomeStruct{
		Email: "email@email.com",
	}

	result := sut.validate.ValidateStruct(invalid)

	assert.Nil(t, result)
}
