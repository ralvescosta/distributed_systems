package errors

type InternalError struct {
	Message string
}

func (e InternalError) Error() string {
	return e.Message
}

func NewInternalError(m string) error {
	return InternalError{Message: m}
}
