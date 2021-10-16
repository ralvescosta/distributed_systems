package errors

import "testing"

func Test_InternalError_Creation(t *testing.T) {
	err := NewInternalError("internal error")

	if err.Error() != "internal error" {
		t.Error("the error message must be the same message when the error was created")
	}
}
