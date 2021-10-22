package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_Conflict_error(t *testing.T) {
	err := NewConflictError("conflict error")

	assert.EqualError(t, err, "conflict error", "the error message must be the same message when the error was created")
}
