# ctrl-c

Graceful shutdown via OS signal handling.

## What's here

`main.go` — first SIGINT cancels a context; second SIGINT hard-exits.

## Todo

- [ ] Add real work inside `run`: process items from a channel, stop accepting new items on cancel, finish the current item before returning.
- [ ] Support both `SIGINT` and `SIGTERM` in `signal.Notify` — production processes need both.
- [ ] Add a shutdown timeout: if `run` hasn't returned 5 s after cancel, force-exit with a log message.
- [ ] Extract the signal wiring into `func WithShutdown(ctx context.Context, run func(context.Context) error) error` so other programs can reuse it without copy-paste.
