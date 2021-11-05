package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IGetBookByIdUseCase interface {
	Perform(ctx context.Context, txn interface{}, id string) (dtos.BookDto, error)
}
