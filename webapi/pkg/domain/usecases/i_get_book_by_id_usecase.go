package usecases

import (
	"context"
	"webapi/pkg/domain/dtos"
)

type IGetBookByIdUseCase interface {
	Perform(ctx context.Context, id string) (dtos.BookDto, error)
}
