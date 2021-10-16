package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"webapi/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zap *zap.Logger
}

func (l Logger) GetHandleFunc() gin.HandlerFunc {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "production" || goEnv == "staging" {
		return l.ProductionLoggerFormater
	}
	return gin.Logger()
}

func (l Logger) Debug(msg string, fields ...interfaces.LogField) {
	l.zap.Debug(msg, convertLogField(fields)...)
}

func (l Logger) Info(msg string, fields ...interfaces.LogField) {
	l.zap.Info(msg, convertLogField(fields)...)
}

func (l Logger) Warn(msg string, fields ...interfaces.LogField) {
	l.zap.Warn(msg, convertLogField(fields)...)
}

func (l Logger) Error(msg string, fields ...interfaces.LogField) {
	l.zap.Error(msg, convertLogField(fields)...)
}

func (l Logger) ProductionLoggerFormater(ctx *gin.Context) {
	startTime := time.Now()
	ctx.Next()
	endTime := time.Now()
	latencyTimeInMileseconds := float64(endTime.Sub(startTime).Nanoseconds() / 1000)

	body, _ := ioutil.ReadAll(ctx.Request.Body)

	l.zap.Info("Request",
		zapcore.Field{
			Key:    "method",
			Type:   zapcore.StringType,
			String: ctx.Request.Method,
		},
		zapcore.Field{
			Key:     "statusCode",
			Type:    zapcore.Int64Type,
			Integer: int64(ctx.Writer.Status()),
		},
		zapcore.Field{
			Key:    "latencyTime",
			Type:   zapcore.StringType,
			String: fmt.Sprintf("%.2f us", latencyTimeInMileseconds),
		},
		zapcore.Field{
			Key:    "clientIP",
			Type:   zapcore.StringType,
			String: ctx.ClientIP(),
		},
		zapcore.Field{
			Key:    "uri",
			Type:   zapcore.StringType,
			String: ctx.Request.RequestURI,
		},
		zapcore.Field{
			Key:    "body",
			Type:   zapcore.StringType,
			String: string(body),
		},
	)
}

func convertLogField(fields []interfaces.LogField) []zap.Field {
	zapFields := []zap.Field{}

	for _, field := range fields {
		zapFields = append(zapFields, zap.Field{
			Key:       field.Key,
			Interface: field.Value,
		})
	}

	return zapFields
}

func configureZapInstance() *zap.Logger {
	goEnv := os.Getenv("GO_ENV")
	logLevel := os.Getenv("LOG_LEVEL")
	var zapLogLevel zapcore.Level
	switch logLevel {
	case "Debug":
		zapLogLevel = zap.DebugLevel
	case "Info":
		zapLogLevel = zap.InfoLevel
	case "Warn":
		zapLogLevel = zap.WarnLevel
	case "Error":
		zapLogLevel = zap.ErrorLevel
	case "Panic":
		zapLogLevel = zap.PanicLevel
	default:
		zapLogLevel = zap.InfoLevel
	}

	var zapInstance *zap.Logger
	switch goEnv {
	case "production":
		zapInstance, _ = zap.NewProduction(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "staging":
		zapInstance, _ = zap.NewProduction(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "development":
		zapInstance, _ = zap.NewDevelopment(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	case "test":
		zapInstance, _ = zap.NewDevelopment(zap.IncreaseLevel(zapLogLevel), zap.AddStacktrace(zap.ErrorLevel))
	default:
		break
	}

	return zapInstance
}

func NewLogger() interfaces.ILogger {
	return Logger{
		zap: configureZapInstance(),
	}
}
