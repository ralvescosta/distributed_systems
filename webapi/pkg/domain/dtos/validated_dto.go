package dtos

type ValidatedDto struct {
	IsValid bool
	Field   string
	Message string
}
