package adapters

import (
	"testing"
)

func Test_Should_Exec_Handler_Successfully(t *testing.T) {
	sut := newHandlerAdaptToTest(false)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 1 {
		t.Error("should called handler once")
	}
}

func Test_Should_Exec_Handler_With_Body_Error(t *testing.T) {
	sut := newHandlerAdaptToTest(true)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}
