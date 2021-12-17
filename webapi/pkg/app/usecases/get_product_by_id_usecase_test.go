package usecases

import (
	"context"
	"testing"
	"webapi/pkg/domain/dtos"

	"github.com/stretchr/testify/assert"
)

func Test_GetProductByIdUC_Should_Execute_Correctly(t *testing.T) {
	config := map[string]mockConfigure{
		"inventoryClient": {
			method:       "GetProductById",
			customResult: dtos.ProductDto{},
			customError:  nil,
		},
	}

	sut := newGetProductByIdUsecaseToTest(config)

	result, err := sut.useCase.Perform(context.Background(), "")

	assert.NoError(t, err)
	assert.IsType(t, result, dtos.ProductDto{})
}
