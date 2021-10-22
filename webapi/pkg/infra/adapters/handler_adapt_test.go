package adapters

import (
	"testing"
)

func Test_Should_Exec_Handler_Successfully(t *testing.T) {
	sut := NewHandlerAdaptToTest(false)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 1 {
		t.Error("should called handler once")
	}
}

func Test_Should_Exec_Handler_With_Body_Error(t *testing.T) {
	sut := NewHandlerAdaptToTest(true)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}
