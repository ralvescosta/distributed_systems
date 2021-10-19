package errors

type BadRequestError struct {
	Message string
}

func (e BadRequestError) Error() string {
	return e.Message
}

func NewBadRequestError(m string) error {
	return BadRequestError{Message: m}
}
