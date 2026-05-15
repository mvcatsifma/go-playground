# iface

Interface composition, type assertions, type switches.

## Key concepts

- **Interface embedding** — compose small focused interfaces rather than defining large ones (`io.Reader + io.Writer → io.ReadWriter`)
- **Implicit satisfaction** — no `implements` keyword; a type satisfies an interface by having the methods
- **Type assertion** — `v, ok := x.(T)`; always use the two-value form to avoid panics
- **Type switch** — `switch v := x.(type)`; the Go way to dispatch on dynamic type
- **Accept interfaces, return structs** — define interfaces in the consumer package, not the producer

## Todo

- [ ] Define `Stringer` and `Sizer`, compose them into `Describer`, implement both on `Circle` and `Rectangle`; write `describe(Describer)` — practice the composition pattern from scratch.
- [ ] Write a function accepting `io.ReadWriter`; test it with `bytes.Buffer` — build the habit of using the narrowest interface that works.
- [ ] Write `sum(vals []any) (float64, error)` using a type switch over `int`, `float64`, `string`; return an error for unsupported types.
- [ ] Refactor a function that took `*os.File` to accept `io.Reader`; verify it now works with `strings.NewReader` in tests — practice the most impactful interface refactor in Go.

## Run

```bash
go run ./iface/
```
