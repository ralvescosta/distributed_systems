package adapters

import (
	"errors"
	"io"
	"testing"
)

func Test_Should_Exec_Handler_Successfully(t *testing.T) {
	sut := NewHandlerAdaptToTest()

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 1 {
		t.Error("should called handler once")
	}
}

func Test_Should_Exec_Handler_With_Body_Error(t *testing.T) {
	sut := NewHandlerAdaptToTest()

	readAllBody = func(r io.Reader) ([]byte, error) {
		return []byte{}, errors.New("Error")
	}

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}
