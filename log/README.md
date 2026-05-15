# log

Standard `log` package: flags, prefix, UTC.

## What's here

`main.go` — configures date, time, UTC, long file, and a per-goroutine prefix.

## Todo

- [ ] Create a dedicated error logger with `log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.LUTC)` and use it alongside the default logger.
- [ ] Write to a `bytes.Buffer` and assert the output format in a test — practice making log output testable.
- [ ] Redirect the default logger to a file with `log.SetOutput(f)` and verify entries survive process restart.
- [ ] Rewrite the same program using `log/slog`; compare how structured fields replace string prefixes.
