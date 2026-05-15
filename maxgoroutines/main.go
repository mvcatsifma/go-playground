package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"time"
)

var cpus = runtime.NumCPU()

func init() {
	log.Println("number of cpu's", cpus)
	runtime.GOMAXPROCS(cpus)
}

// Spawns one worker goroutine per CPU; each drains itemChan until it is closed.
// SIGINT stops the producer and closes the channel, triggering a clean shutdown via WaitGroup.
func main() {
	itemChan := make(chan int)

	var wg sync.WaitGroup
	wg.Add(cpus)

	for j := 0; j < cpus; j++ {
		go func() {
			for {
				id, ok := <-itemChan
				if !ok {
					fmt.Println("itemChan closed")
					wg.Done()
					return
				}
				fmt.Println("process item ", id)
				time.Sleep(1 * time.Second)
			}
		}()
	}

	terminateChan := make(chan struct{})
	go func() {
	WriteLoop:
		for i := 0; i < 100; i++ {
			select {
			case <-terminateChan:
				break WriteLoop
			default:
				itemChan <- i
			}
		}
	}()

	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)
readInterruptLoop:
	for {
		select {
		case sig := <-osInterruptChannel:
			switch sig {
			case os.Interrupt: // SIGIINT received
				log.Println("SIGIINT received, shutting down")

				break readInterruptLoop
			default:
				log.Printf("unknown signal received: %v", sig)
			}
		}
	}

	// terminate the write loop
	terminateChan <- struct{}{}
	close(itemChan)
	wg.Wait()
}
