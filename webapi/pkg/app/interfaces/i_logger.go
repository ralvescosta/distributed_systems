package interfaces

import "github.com/gin-gonic/gin"

type LogField struct {
	Key   string
	Value interface{}
}

type ILogger interface {
	GetHandleFunc() gin.HandlerFunc
	Debug(msg string, fields ...LogField)
	Info(msg string, fields ...LogField)
	Warn(msg string, fields ...LogField)
	Error(msg string, fields ...LogField)
}
