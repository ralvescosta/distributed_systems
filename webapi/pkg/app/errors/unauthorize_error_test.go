package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Create_Unauthorize_error(t *testing.T) {
	err := NewConflictError("unauthorize error")

	assert.EqualError(t, err, "unauthorize error", "the error message must be the same message when the error was created")
}
