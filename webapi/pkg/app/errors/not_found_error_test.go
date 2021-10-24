package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_NotFound_error(t *testing.T) {
	err := NewNotFoundError("notFound error")

	assert.EqualError(t, err, "notFound error", "the error message must be the same message when the error was created")
}
