package ctxlog

import "context"

func (l *CtxLogger) Error(ctx context.Context, msg string) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Error(msg)
}
