package utils

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	loc, _ := time.LoadLocation("UTC")
	t = t.In(loc)

	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func GetLog() *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	cfg.Development = true
	cfg.DisableCaller = false
	cfg.DisableStacktrace = false
	cfg.Encoding = "console" // "console" or "json"
	cfg.EncoderConfig.EncodeTime = SyslogTimeEncoder
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.Level.SetLevel(zap.DebugLevel)

	logger, err := cfg.Build()
	if err != nil {
		log.Panicf("Could not build logger, err: %v", err)
	}
	defer logger.Sync() // Flushes buffer, if any
	log := logger.Sugar()

	return log
}
