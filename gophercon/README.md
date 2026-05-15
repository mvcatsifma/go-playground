# gophercon

Channel fan-out and select-based shutdown via `close(quit)`.

## What's here

`main.go` — producer sends messages on `c`, closes `quit` when done; main selects on both.

## Todo

- [ ] Replace `os.Exit(0)` with a clean return so the program can be tested without killing the process.
- [ ] Add a `time.After(1 * time.Second)` case to the select — always add a timeout case in production select loops.
- [ ] Convert `quit` to a `context.Context` cancel and rewrite the shutdown; compare the two patterns side by side.
- [ ] Write a test that verifies all messages arrive and the program exits without a goroutine leak.
