package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
	"webapi/pkg/domain/usecases"
)

type getBookByIdUseCase struct{}

func (getBookByIdUseCase) Perform(ctx context.Context, txn interface{}, id string) (dtos.BookDto, error) {
	return dtos.BookDto{}, nil
}

func NewGetBookByIdUseCase() usecases.IGetBookByIdUseCase {
	return getBookByIdUseCase{}
}
