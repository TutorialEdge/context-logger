package ctxlog_test

import (
	"context"
	"testing"

	"github.com/TutorialEdge/ctxlog"
)

func TestLogger(t *testing.T) {
	t.Run("test logger instantiates properly", func(t *testing.T) {
		ctx := context.Background()
		log := ctxlog.New()
		ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
			"trace_id": "my-trace-id",
		})

		log.Info(ctx, "hello world")
	})

	t.Run("test logger WithJSONFormat", func(t *testing.T) {
		ctx := context.Background()
		log := ctxlog.New(
			ctxlog.WithJSONFormat(),
		)
		ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
			"trace_id": "my-trace-id",
		})

		log.Info(ctx, "hello world")
		log.Info(ctx, "something else")
	})
}
