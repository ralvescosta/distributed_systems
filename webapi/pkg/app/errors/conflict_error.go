package errors

type ConflictError struct {
	Message string
}

func (e ConflictError) Error() string {
	return e.Message
}

func NewConflictError(m string) error {
	return ConflictError{Message: m}
}
