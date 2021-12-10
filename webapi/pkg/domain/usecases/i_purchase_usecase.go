package usecases

import "context"

type IPurchaseUseCase interface {
	Perform(ctx context.Context) error
}
