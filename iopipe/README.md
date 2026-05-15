# iopipe

`io.Pipe` — synchronous in-memory pipe; no intermediate buffer.

## Key concepts

- **`io.Pipe()`** — returns `*PipeReader` and `*PipeWriter`; writes block until the reader consumes
- **`pw.CloseWithError(err)`** — propagates an error to the reader
- **`pr.CloseWithError(err)`** — signals the writer the reader is done
- **Use case** — chain an encoder/compressor to a consumer without buffering the whole payload in memory

## Todo

- [ ] Write the simplest pipe: send a string from a goroutine, read it in main — build the baseline mental model.
- [ ] Wrap the writer end in `gzip.NewWriter` and the reader end in `gzip.NewReader`; verify the decompressed output matches the input — the most common production use.
- [ ] Have the writer fail mid-stream with `pw.CloseWithError(err)`; verify the reader receives the error from `Read`.
- [ ] Combine `io.Pipe` with `io.TeeReader` to simultaneously stream data to a consumer and accumulate a SHA-256 hash of the bytes — practice composing `io.Reader` values.

## Run

```bash
go run ./iopipe/
```
