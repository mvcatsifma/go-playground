# options

Functional options with undo/restore.

## What's here

`main.go` — applies `Verbosity` and defers restoring the previous value.  
`entity/foo.go` — `Option` sets an option and returns the previous state as another option.

## Todo

- [ ] Add `Timeout(d time.Duration) option`; apply both `Verbosity` and `Timeout` in sequence and verify both restore on defer.
- [ ] Rewrite `Foo` using the more common variadic `func(o *options)` style; compare the API ergonomics and decide which you prefer.
- [ ] Write a test: apply an option inside a helper, defer restore, assert the value reverts when the helper returns.
- [ ] Add `Reset() option` that zeroes all fields; verify it composes correctly with other options in a chain.
