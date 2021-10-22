package logger

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Should_Execute_GetHandleFunc_Correctly(t *testing.T) {
	sut := NewLoggerSpyToTest()

	result := reflect.TypeOf(sut.logger.GetHandleFunc()).Name()

	assert.Equal(t, result, "HandlerFunc", "Should return HandlerFunc")
}

func Test_Should_Execute_Logger_Methods_Correctly(t *testing.T) {
	sut := NewLoggerSpyToTest()

	sut.logger.Debug("some message")
	sut.logger.Info("some message")
	sut.logger.Warn("some message")
	sut.logger.Error("some message")
}
