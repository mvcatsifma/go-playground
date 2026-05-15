package main

import (
	"fmt"
	"log"
)

type A struct {
}

func (a A) String() string {
	return "A"
}

func main() {
	a := &A{}
	fn(a)
}

// fn accepts any T that satisfies fmt.Stringer, logs it, and returns it unchanged.
func fn[T fmt.Stringer](t T) T {
	log.Println(t.String())
	return t
}
