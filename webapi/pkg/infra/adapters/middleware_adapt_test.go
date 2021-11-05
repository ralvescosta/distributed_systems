package adapters

import (
	"net/http"
	"testing"
)

func Test_MidAdapt_Should_Exec_Middleware_Successfully(t *testing.T) {
	sut := newMiddlewareAdaptToTest(false, http.StatusOK)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 1 {
		t.Error("should called handler once")
	}
}

func Test_MidAdapt_Should_Exec_Middleware_With_Body_Error(t *testing.T) {
	sut := newMiddlewareAdaptToTest(true, http.StatusOK)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}

func Test_MidAdapt_Should_Abort_If_Handler_Return_Error(t *testing.T) {
	sut := newMiddlewareAdaptToTest(true, http.StatusBadRequest)

	sut.adapt(sut.ctx)

	if *sut.handlerCalledTimes != 0 {
		t.Error("Shouldn't call handler when body is unformatted")
	}
}
