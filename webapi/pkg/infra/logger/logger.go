package logger

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"webapi/pkg/app/interfaces"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	*zap.Logger
}

func (l logger) GetHandleFunc() gin.HandlerFunc {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "production" || goEnv == "staging" {
		return l.HttpRequestLogger
	}
	return gin.Logger()
}

func (l logger) HttpRequestLogger(ctx *gin.Context) {
	startTime := time.Now()
	ctx.Next()
	endTime := time.Now()
	latencyTimeInMileseconds := float64(endTime.Sub(startTime).Nanoseconds() / 1000)

	body, _ := ioutil.ReadAll(ctx.Request.Body)

	l.Info("Request",
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

func createLoggerInstance() *zap.Logger {
	return zap.New(
		configureZapCore(),
		zap.IncreaseLevel(getIncreaseLevel()),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}

func configureZapCore() zapcore.Core {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	if os.Getenv("GO_ENV") == "development" {
		debugging := zapcore.Lock(os.Stdout)
		errors := zapcore.Lock(os.Stderr)
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

		return zapcore.NewTee(
			zapcore.NewCore(consoleEncoder, debugging, lowPriority),
			zapcore.NewCore(consoleEncoder, errors, highPriority),
		)
	}

	logIo, err := os.Create(getLogPath())
	if err != nil {
		err = fmt.Errorf("server.Start - create log writer")
		log.Fatal(err)
	}

	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	return zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, logIo, lowPriority),
		zapcore.NewCore(jsonEncoder, logIo, highPriority),
	)
}

func getLogPath() string {
	logPath := os.Getenv("LOG_PATH")
	if logPath != "" {
		return logPath
	}

	return "file.log"
}

func getIncreaseLevel() zapcore.Level {
	logLevel := os.Getenv("LOG_LEVEL")
	switch logLevel {
	case "Debug":
		return zap.DebugLevel
	case "Info":
		return zap.InfoLevel
	case "Warn":
		return zap.WarnLevel
	case "Error":
		return zap.ErrorLevel
	case "Panic":
		return zap.PanicLevel
	default:
		return zap.InfoLevel
	}
}

func NewLogger() interfaces.ILogger {
	return logger{
		createLoggerInstance(),
	}
}
