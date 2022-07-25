ctx-logger
===========


```go
ctx := context.Background()
		log := ctxlog.New(
			ctxlog.WithJSONFormat(),
		)
		ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
			"trace_id": "my-trace-id",
		})

		log.Info(ctx, "hello world")
```