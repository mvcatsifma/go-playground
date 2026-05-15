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

func fn[T fmt.Stringer](t T) T {
	log.Println(t.String())
	return t
}
