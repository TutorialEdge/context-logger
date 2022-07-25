package ctxlog_test

import (
	"context"
	"github.com/TutorialEdge/ctxlog"
	"testing"
)

func TestErrorLogs(t *testing.T) {
	t.Run("test error log output", func(t *testing.T) {
		ctx := context.Background()
		log := ctxlog.New()
		ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
			"trace_id": "my-trace-id",
		})

		log.Error(ctx, "some error message")
	})
}
