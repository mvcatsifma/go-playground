package main

import (
	"dsen.nl/go-playground/options/entity"
	"fmt"
)

func main() {
	foo := &entity.Foo{}
	prev := foo.Option(entity.Verbosity(1))
	defer foo.Option(prev)
	fmt.Printf("%+v\n", foo)
}
