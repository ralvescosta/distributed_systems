package adapters

import (
	"net/http"
	"testing"
)

func Test_HandleAdapt_Should_Exec_Handler_Successfully(t *testing.T) {
	sut := newHandlerAdaptToTest(false, http.StatusOK)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 1 {
		t.Error("should called handler once")
	}
}

func Test_HandleAdapt_Should_Exec_Handler_With_Body_Error(t *testing.T) {
	sut := newHandlerAdaptToTest(true, http.StatusOK)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}
func Test_HandleAdapt_Should_Abort_If_Handler_Return_Error(t *testing.T) {
	sut := newMiddlewareAdaptToTest(true, http.StatusBadRequest)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}
