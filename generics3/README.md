# generics3

Generic functions constrained to an interface.

## What's here

`main.go` — `fn[T fmt.Stringer](t T) T` logs any `Stringer` and returns it unchanged.

## Todo

- [ ] Write `Transform[T fmt.Stringer](items []T, f func(T) T) []T`; apply it to a slice of `A` values.
- [ ] Define a composite constraint `interface { fmt.Stringer; comparable }` and write `Deduplicate[T]` that removes consecutive equal values.
- [ ] Try passing `A` (value receiver) vs `*A` (pointer receiver) to `fn`; observe the compiler error — understanding method sets is essential.
- [ ] Write `StringsOf[T fmt.Stringer](items []T) []string` that collects `String()` output into a slice; use it in a test assertion.
