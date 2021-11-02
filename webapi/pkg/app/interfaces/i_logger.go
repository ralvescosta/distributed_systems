package interfaces

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LogField struct {
	Key   string
	Value interface{}
}

type ILogger interface {
	GetHandleFunc() gin.HandlerFunc
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}
