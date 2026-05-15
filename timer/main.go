package main

import (
	"fmt"
	"time"
)

// A negative duration creates a timer that fires immediately on the first read.
func main() {
	timer := time.NewTimer(-1 * time.Second)

	select {
	case <-timer.C:
		fmt.Println("timer expired")
	}
}
