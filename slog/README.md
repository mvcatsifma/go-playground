# slog

Structured logging with `log/slog` (Go 1.21+).

## Key concepts

- **`slog.New(handler)`** — `TextHandler` for humans, `JSONHandler` for log aggregators
- **`logger.With(k, v)`** — child logger with pre-attached attrs (e.g. request ID, trace ID)
- **`logger.WithGroup(name)`** — namespace attrs under a key (`"db.host"`, `"http.status"`)
- **`logger.LogAttrs`** — typed `slog.Attr` values; avoids allocations in hot paths
- **`slog.SetDefault`** — replaces the process-wide default used by `slog.Info` etc.

## Todo

- [ ] Log the same message with `TextHandler` and `JSONHandler` side by side — know which to reach for and why.
- [ ] Build a request-scoped logger: `logger.With("trace_id", id).WithGroup("http")` then log method, path, and status; verify all attrs appear nested correctly.
- [ ] Configure a handler with `HandlerOptions{Level: slog.LevelWarn}` and verify `Info` calls are silenced — this is how you gate log verbosity per environment.
- [ ] Implement a custom `slog.Handler` that writes coloured output; wire it with `slog.SetDefault` — practice the Handler interface so you can adapt slog to any sink.

## Run

```bash
go run ./slog/
```
