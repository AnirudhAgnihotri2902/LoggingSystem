package logger

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field = zap.Field

// Logger is the interface to allow common logging functionalities.
type Logger interface {
	InfoContext(ctx context.Context, msg string, fields ...Field)
	DebugContext(ctx context.Context, msg string, fields ...Field)
	ErrorContext(ctx context.Context, msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Debug(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	//With(key string, value string)
	Warn(msg string, fields ...Field)
}

type LoggerImpl struct {
	logger *otelzap.Logger
}

// addContextFieldsToLogs appends context fields such as request_id and trace_id to logs.
func addContextFieldsToLogs(ctx context.Context, fields ...Field) []Field {
	fields = append(fields, zap.Any("request_id", ctx.Value("requestID")))
	fields = append(fields, zap.Any("user_id", ctx.Value("userID")))
	//fields = append(fields, zap.Any("request_id", ctx.Value("trace_id")))
	return fields
}

// InfoContext logs the message along with fields from context on info level.
func (l *LoggerImpl) InfoContext(ctx context.Context, msg string, fields ...Field) {
	fields = addContextFieldsToLogs(ctx, fields...)
	l.logger.Ctx(ctx).Info(msg, fields...)
}

// DebugContext logs the message along with fields from context on debug level.
func (l *LoggerImpl) DebugContext(ctx context.Context, msg string, fields ...Field) {
	fields = addContextFieldsToLogs(ctx, fields...)
	l.logger.Ctx(ctx).Debug(msg, fields...)
}

// ErrorContext logs the message along with fields from context on Error level.
func (l *LoggerImpl) ErrorContext(ctx context.Context, msg string, fields ...Field) {
	fields = addContextFieldsToLogs(ctx, fields...)
	l.logger.Ctx(ctx).Error(msg, fields...)
}

// Info logs the message on info level.
func (l *LoggerImpl) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

// Debug logs the message on debug level.
func (l *LoggerImpl) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

// Error logs the message on Error level.
func (l *LoggerImpl) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

//// With logs the message on Error level.
//func (l *LoggerImpl) With(key string, value string) {
//	l.logger = l.logger.With(zap.String(key, value))
//}

// Warn logs the message on Warn level.
func (l *LoggerImpl) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

// Int return a integer value log field.
func Int(key string, val int) Field {
	return zap.Int(key, val)
}

// Int64 return a integer value log field.
func Int64(key string, val int64) Field {
	return zap.Int64(key, val)
}

// String returns a string value log field.
func String(key string, val string) Field {
	return zap.String(key, val)
}

// Float64 returns a float value log field.
func Float64(key string, val float64) Field {
	return zap.Float64(key, val)
}

// Bool returns a bool value log field.
func Bool(key string, val bool) Field {
	return zap.Bool(key, val)
}

// Error return a error value log field.
func Error(val error) Field {
	return zap.Error(val)
}

// Any returns a error value with generic data field.
func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}

// New creates a new zap logger with the given log level, service name and environment.
func New(level zapcore.Level, serviceName, environment string) (Logger, error) {
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

	return &LoggerImpl{
		logger: otelzap.New(logger.WithOptions(zap.AddCallerSkip(1)),
			otelzap.WithMinLevel(zap.InfoLevel),
			//otelzap.WithExtraFields(zapcore.Field{Key: "user_id", Type: zapcore.StringType}),
			otelzap.WithCallerDepth(1),
			otelzap.WithTraceIDField(true),
		),
	}, nil
}

func NewExample() Logger {
	logger := zap.NewExample()
	return &LoggerImpl{
		logger: otelzap.New(logger.WithOptions(zap.AddCallerSkip(1)),
			otelzap.WithMinLevel(zap.InfoLevel)),
	}
}
