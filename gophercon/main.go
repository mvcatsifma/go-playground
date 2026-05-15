package main

import (
	"fmt"
	"os"
)

func main() {
	c := make(chan string)
	quit := make(chan struct{})

	// Producer closes quit when all messages are sent; the select loop exits on that signal.
	go func(messages []string) {
		for _, s := range messages {
			c <- s
		}
		close(quit)
	}([]string{"hi", "bye"})

	for {
		select {
		case message := <-c:
			fmt.Println(message)
		case <-quit:
			fmt.Println("shutting down")
			os.Exit(0)
		}
	}
}
