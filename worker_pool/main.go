package main

import (
	"fmt"
	"math/rand"
	"time"
)

type token struct {
}

var hugeSlice []int
func init() {
	for i := 0; i < 20; i++ {
		hugeSlice = append(hugeSlice, i)
	}
}

func main() {
	const limit = 2
	sem := make(chan token, limit)

	// hugeSlice could also be a channel
	for _, task := range hugeSlice {
		sem <- token{}
		go func(task int) {
			fmt.Printf("executing task %v\n", task)
			time.Sleep(time.Duration(rand.Int63n(2)) * time.Second)
			<- sem
		}(task)
	}

	fmt.Println("all tasks started, wait for done")
	for n := limit; n > 0; n-- {
		sem <- token{}
	}

	fmt.Println("done")
}
