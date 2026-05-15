package main

import (
	"fmt"
)

func main() {
	foo := HelloWriter(Hello)
	_, _ = foo.Write([]byte("hello there"))
}

// HelloWriter adapts a plain function to the io.Writer interface.
type HelloWriter func(str string)

func (f HelloWriter) Write(p []byte) (int, error) {
	f(string(p))
	return 0, nil
}

func Hello(str string) {
	fmt.Println(str)
}
