# generics

Type parameters, inline union constraints, named constraints.

## What's here

`main.go` — `SumIntsOrFloats` (inline union), `SumNumbers` (named `Number` constraint), side-by-side with non-generic versions.

## Todo

- [ ] Write `Map[In, Out any](s []In, f func(In) Out) []Out` and `Filter[V any](s []V, keep func(V) bool) []V`; chain them on a real slice.
- [ ] Write `Min[T cmp.Ordered](a, b T) T` and `Max[T cmp.Ordered](a, b T) T` — practice using stdlib constraints rather than defining your own.
- [ ] Write a generic `Stack[T any]` with `Push`, `Pop`, and `Peek`; test it with both `string` and `int`.
- [ ] Extend `Number` to include `int` and `float32`; update all callers — practice the mechanical work of changing a constraint.
