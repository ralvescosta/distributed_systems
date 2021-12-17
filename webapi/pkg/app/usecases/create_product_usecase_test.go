package usecases

import (
	"context"
	"testing"
	"webapi/pkg/domain/dtos"

	"github.com/stretchr/testify/assert"
)

func Test_CreateProductUC_Should_Execute_Correctly(t *testing.T) {
	config := map[string]mockConfigure{
		"inventoryClient": {
			method:       "RegisterProduct",
			customResult: dtos.ProductDto{},
			customError:  nil,
		},
	}

	sut := newCreateProductUsecaseToTest(config)

	result, err := sut.useCase.Perform(context.Background(), dtos.ProductDto{})

	assert.NoError(t, err)
	assert.IsType(t, result, dtos.ProductDto{})
}
