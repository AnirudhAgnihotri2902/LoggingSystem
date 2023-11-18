package logger

import (
	"context"

	"go.uber.org/zap"
)

// ZapContext create context with logger.
func ZapContext(ctx context.Context, logger *zap.Logger, requestID string) context.Context {
	ctx = ToContext(ctx, logger)
	AddFields(ctx, zap.String("request_id", requestID))
	return ctx
}

// FromContext extracts logger from context.
func FromContext(ctx context.Context) *zap.Logger {
	return Extract(ctx)
}
