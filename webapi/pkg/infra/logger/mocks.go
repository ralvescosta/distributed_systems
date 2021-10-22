package logger

import (
	"webapi/pkg/app/interfaces"

	"go.uber.org/zap"
)

type LoggerSpyToTest struct {
	logger interfaces.ILogger
}

func NewLoggerSpyToTest() LoggerSpyToTest {
	return LoggerSpyToTest{
		logger: NewLoggerSpy(),
	}
}

type LoggerToTest struct {
	zapInstance *zap.Logger
	logger      interfaces.ILogger
}

func NewLoggerToTest() LoggerToTest {
	zap, _ := zap.NewDevelopment(zap.IncreaseLevel(zap.DebugLevel), zap.AddStacktrace(zap.ErrorLevel))

	sut := &Logger{
		zap: zap,
	}

	return LoggerToTest{
		zapInstance: zap,
		logger:      sut,
	}
}
