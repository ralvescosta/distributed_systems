package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"webapi/pkg/app/interfaces"
	"webapi/pkg/domain/dtos"
)

type _validator struct{}

func (_validator) ValidateStruct(m interface{}) []dtos.ValidatedDto {
	v := validator.New()
	err := v.Struct(m)

	if err == nil {
		return nil
	}

	validationErrors := err.(validator.ValidationErrors)

	var validatedErros []dtos.ValidatedDto
	for _, validationErr := range validationErrors {
		validatedErros = append(validatedErros, dtos.ValidatedDto{
			IsValid: false,
			Field:   validationErr.Field(),
			Message: fmt.Sprintf("%s is invalid", validationErr.Field()),
		})
	}

	return validatedErros
}

func NewValidator() interfaces.IValidator {
	return _validator{}
}
