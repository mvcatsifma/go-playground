# httphandler

Function-as-`io.Writer` — adapting a plain function to satisfy an interface.

## What's here

`main.go` — `HelloWriter` is a function type with a `Write` method.

## Todo

- [ ] Wire `HelloWriter` into an HTTP handler: use it as the `io.Writer` for `json.NewEncoder` and write a JSON response.
- [ ] Implement `LineWriter` — a function type that buffers bytes until `\n` then calls the underlying function with the complete line.
- [ ] Write `CountingWriter` wrapping any `io.Writer` that counts bytes written; assert the count after `fmt.Fprintf`.
- [ ] Fan out to two functions simultaneously using `io.MultiWriter(hw1, hw2)` — practice composing `io.Writer` values.
