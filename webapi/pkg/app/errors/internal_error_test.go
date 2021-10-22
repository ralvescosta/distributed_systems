package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_InternalError(t *testing.T) {
	err := NewInternalError("internal error")

	assert.EqualError(t, err, "internal error", "the error message must be the same message when the error was created")
}
