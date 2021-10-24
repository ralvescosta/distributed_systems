package interfaces

import "webapi/pkg/domain/dtos"

type IValidator interface {
	ValidateStruct(m interface{}) []dtos.ValidatedDto
}
