package errors

type NotFoundError struct {
	Message string
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewNotFoundError(m string) error {
	return NotFoundError{Message: m}
}
