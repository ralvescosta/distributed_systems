package logger

import (
	"webapi/pkg/app/interfaces"

	"go.uber.org/zap"
)

type loggerSpyToTest struct {
	logger interfaces.ILogger
}

func newLoggerSpyToTest() loggerSpyToTest {
	return loggerSpyToTest{
		logger: NewLoggerSpy(),
	}
}

type loggerToTest struct {
	zapInstance *zap.Logger
	logger      interfaces.ILogger
}

func newLoggerToTest() loggerToTest {
	zap, _ := zap.NewDevelopment(zap.IncreaseLevel(zap.DebugLevel), zap.AddStacktrace(zap.ErrorLevel))

	sut := logger{
		zap,
	}

	return loggerToTest{
		zapInstance: zap,
		logger:      sut,
	}
}
