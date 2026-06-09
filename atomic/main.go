package atomic

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

// runCleanExit demonstrates using atomic.Bool as a shutdown flag for graceful termination.
// Workers check the flag before starting work; signal handler sets it to prevent new work.
// In-progress work completes naturally before exit.
func runCleanExit() {
	wg := &sync.WaitGroup{}
	var stop atomic.Bool // Shutdown flag - thread-safe without mutex

	// Spawn 20 workers that simulate 5-second work cycles
	for i := 0; i < 20; i++ {
		wg.Go(func() {
			if stop.Load() { // Atomic read - check if shutdown requested
				return // Don't start new work if stopping
			}
			// Simulate work (e.g., processing a batch, completing a transaction)
			time.Sleep(5 * time.Second)
		})
	}

	// Register signal handler for Ctrl+C
	signalChan := make(chan os.Signal, 1) // Buffered to prevent missing signals
	signal.Notify(signalChan, os.Interrupt)

	fmt.Println("Workers running. Press Ctrl+C for graceful shutdown...")
	<-signalChan // Block until signal received
	fmt.Println("\nShutdown signal received. Waiting for in-progress work to complete...")

	stop.Store(true) // Atomic write - signal workers to stop accepting new work

	wg.Wait() // Wait for all in-progress work to finish

	fmt.Println("All workers completed. Clean exit.")
}

// concurrentIncrement spawns nr goroutines that each increment a shared
// atomic counter exactly once. Returns the final count after all goroutines
// complete. With correct atomic operations, the result will always equal nr.
// Run with -race to verify no data races occur.
func concurrentIncrement(nr int) int32 {
	wg := &sync.WaitGroup{}
	var count atomic.Int32

	for i := 0; i < nr; i++ {
		wg.Go(func() { // Go 1.26+: automatically handles Add(1)/Done()
			count.Add(1) // Atomic increment - thread-safe, no mutex needed
		})
	}
	wg.Wait() // Block until all goroutines complete

	return count.Load() // Atomic read of final value
}
