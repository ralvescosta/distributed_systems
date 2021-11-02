package logger

import (
	"webapi/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type loggerSpy struct {
	GetHandleFuncCallerCount int
	DebugCallerCount         int
	InfoCallerCount          int
	WarnCallerCount          int
	ErrorCallerCount         int
	FatalCallerCount         int
}

func (l *loggerSpy) GetHandleFunc() gin.HandlerFunc {
	l.GetHandleFuncCallerCount++
	return func(*gin.Context) {}
}
func (l *loggerSpy) Debug(msg string, fields ...zap.Field) {}
func (l *loggerSpy) Info(msg string, fields ...zap.Field)  {}
func (l *loggerSpy) Warn(msg string, fields ...zap.Field)  {}
func (l *loggerSpy) Error(msg string, fields ...zap.Field) {}
func (l *loggerSpy) Fatal(msg string, fields ...zap.Field) {}

func NewLoggerSpy() interfaces.ILogger {
	return &loggerSpy{
		GetHandleFuncCallerCount: 0,
		DebugCallerCount:         0,
		InfoCallerCount:          0,
		WarnCallerCount:          0,
		ErrorCallerCount:         0,
		FatalCallerCount:         0,
	}
}
