package entity

type option func(*Foo) option

type Foo struct {
	Verbosity int
}

// Option sets the options specified.
// It returns an option to restore the last arg's previous value.
func (f *Foo) Option(opts ...option) (previous option) {
	for _, opt := range opts {
		previous = opt(f)
	}
	return previous
}

func Verbosity(i int) option {
	return func(foo *Foo) option {
		previous := foo.Verbosity
		foo.Verbosity = i
		return Verbosity(previous)
	}
}
