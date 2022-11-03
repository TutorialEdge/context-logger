package ctxlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (l *CtxLogger) Warn(ctx context.Context, msg string) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Warn(msg)
}

func (l *CtxLogger) Warn(ctx context.Context, logFunc logrus.LogFunction) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).WarnFn(logFunc)
}
