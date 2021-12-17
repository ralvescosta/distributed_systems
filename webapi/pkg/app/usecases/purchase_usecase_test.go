package usecases

import (
	"context"
	"testing"
	"webapi/pkg/domain/dtos"

	"github.com/stretchr/testify/assert"
)

func Test_PrucaseUC_Should_Execute_Correctly(t *testing.T) {
	config := map[string]mockConfigure{
		"messageBroker": {
			method:      "Publisher",
			customError: nil,
		},
	}

	sut := newCreatePurchaseUsecaseTest(config)

	err := sut.useCase.Perform(context.Background(), dtos.CreatePurchaseDto{})

	assert.NoError(t, err)
}
