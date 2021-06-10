package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig
	l, err := config.Build(zap.AddCallerSkip(1))

	//l, err := zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	log = l
}

func Info(m string, fields ...zap.Field) {
	log.Info(m, fields...)
}

func Debug(m string, fields ...zap.Field) {
	log.Debug(m, fields...)
}

func Error(m string, fields ...zap.Field) {
	log.Error(m, fields...)
}

func Fatal(m string, fields ...zap.Field) {
	log.Fatal(m, fields...)
}
