package logger

import (
	"context"
	. "my-app-tx/utils/constants"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var baseLogger *zap.Logger

func InitLogger() {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:      "time",
		LevelKey:     "level",
		CallerKey:    "caller",
		MessageKey:   "msg",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		zap.DebugLevel,
	)

	baseLogger = zap.New(core, zap.AddCaller())
}

// Logger con contexto (similar a log4j MDC)
func WithContext(ctx context.Context) *zap.Logger {
	tenantID, _ := ctx.Value(HEADER_TENANT_ID).(string)
	requestID, _ := ctx.Value(HEADER_REQUEST_ID).(string)

	return baseLogger.With(
		zap.String("tenant", tenantID),
		zap.String("request_id", requestID),
	)
}
