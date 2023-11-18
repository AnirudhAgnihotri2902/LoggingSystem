package log

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerDeprecated interface {
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}

// NewDeprecatedLogger creates a new zap logger with the given log level, service name and environment.
func NewDeprecatedLogger(level zapcore.Level, serviceName, environment string) (LoggerDeprecated, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(level)
	config.DisableStacktrace = true
	config.Sampling = nil
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.InitialFields = map[string]interface{}{
		"service": serviceName,
		"env":     environment,
	}
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return NewOpenTelemetryLogger(logger), nil
}

func NewOpenTelemetryLogger(logger *zap.Logger) LoggerDeprecated {
	return otelzap.New(logger)
}

func LoggerWithCtx(ctx context.Context, logger LoggerDeprecated) LoggerDeprecated {
	zapLogger := logger.(*otelzap.Logger)
	return zapLogger.Ctx(ctx)
}

// NewDiscard creates logger which output to ioutil.Discard.
// This can be used for testing.
func NewDiscard() LoggerDeprecated {
	return zap.NewNop()
}

func Level(level string) (zapcore.Level, error) {
	level = strings.ToUpper(level)
	var l zapcore.Level
	switch level {
	case "DEBUG":
		l = zapcore.DebugLevel
	case "INFO":
		l = zapcore.InfoLevel
	case "ERROR":
		l = zapcore.ErrorLevel
	default:
		return l, fmt.Errorf("invalid loglevel: %s", level)
	}
	return l, nil
}

// CodeToLevel decides log level based on return grpc code.
func CodeToLevel(code int64) zapcore.Level {
	switch code {
	case http.StatusOK:
		return zap.DebugLevel
	case http.StatusNotFound:
		return zap.DebugLevel
	case http.StatusInternalServerError:
		return zap.ErrorLevel
	case http.StatusBadRequest:
		return zap.InfoLevel
	default:
		return zap.InfoLevel
	}
}
