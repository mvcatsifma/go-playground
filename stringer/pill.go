//go:generate stringer -type=Pill -linecomment

package stringer

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol   // Foo
	Acetaminophen = Paracetamol
)
