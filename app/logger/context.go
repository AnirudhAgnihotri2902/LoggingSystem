package log

import (
	"context"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"go.uber.org/zap"
)

type ctxMarker struct{}

var ctxMarkerKey = &ctxMarker{}

// AddLoggerToContext create context with logger.
func AddLoggerToContext(ctx context.Context, logger Logger, requestID string) context.Context {
	var zapLogger *LoggerImpl
	var ok bool
	zapLogger, ok = logger.(*LoggerImpl)
	if !ok || zapLogger == nil {
		// fix here with default logger.
		zapLogger = &LoggerImpl{logger: otelzap.New(zap.NewExample())}
	}
	ctx = ctxzap.ToContext(ctx, zapLogger.logger.Logger)
	ctxzap.AddFields(ctx, zap.String("request_id", requestID))
	return ctx
}

// FromContext extracts logger from context.
func FromContext(ctx context.Context) Logger {
	logger := zap.NewExample().With(ctxzap.TagsToFields(ctx)...)
	return &LoggerImpl{
		logger: otelzap.New(logger),
	}
}
