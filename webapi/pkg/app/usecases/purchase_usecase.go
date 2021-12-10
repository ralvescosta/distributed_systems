package usecases

import (
	"context"
	"webapi/pkg/domain/usecases"
)

type purchaseUseCase struct{}

func (purchaseUseCase) Perform(ctx context.Context) error {
	return nil
}

func NewPruchaseUseCase() usecases.IPurchaseUseCase {
	return purchaseUseCase{}
}
