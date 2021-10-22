package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_BadRequest_error(t *testing.T) {
	err := NewConflictError("badRequest error")

	assert.EqualError(t, err, "badRequest error", "the error message must be the same message when the error was created")
}
