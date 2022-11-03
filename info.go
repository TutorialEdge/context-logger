package ctxlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

func (l *CtxLogger) Info(ctx context.Context, msg string) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).Info(msg)
}

func (l *CtxLogger) InfoFn(ctx context.Context, logFun logrus.LogFunction) {
	fields := convertFieldsToLogrusFields(getFields(ctx))
	l.Log.WithFields(fields).InfoFn(logFun)
}
