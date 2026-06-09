package atomic

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	version int
}

// updateAndReadConfig demonstrates the read-copy-update (RCU) pattern with atomic.Pointer.
// Multiple readers continuously load config without locks while a single writer periodically
// updates it. Readers always see consistent snapshots - never partial updates.
func updateAndReadConfig() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	var stop atomic.Bool
	var ap atomic.Pointer[Config] // Atomic config pointer - enables lock-free RCU
	ap.Store(&Config{})           // Initialize with version 0

	var wg sync.WaitGroup

	// Spawn 10 reader goroutines - continuously read config every 2 seconds
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			for !stop.Load() { // Poll until shutdown
				time.Sleep(2 * time.Second)
				cfg := ap.Load() // Atomic read - no lock, always fast
				fmt.Printf("Reader #%d: config version %d\n", i, cfg.version)
			}
			fmt.Printf("Reader #%d stopped\n", i)
		})
	}

	// Spawn single writer goroutine - updates config every 5 seconds
	wg.Go(func() {
		for !stop.Load() { // Poll until shutdown
			time.Sleep(5 * time.Second)
			old := ap.Load()    // Read current config
			updated := &Config{ // Copy and modify (never mutate in-place)
				version: old.version + 1,
			}
			ap.Store(updated) // Atomic pointer swap - readers see new version instantly
			fmt.Printf("Manager: updated config to version %d\n", updated.version)
		}
		fmt.Println("Manager stopped")
	})

	fmt.Println("RCU config demo: 10 readers (every 2s), 1 writer (every 5s)")
	fmt.Println("Press Ctrl+C to stop...")
	<-signalChan // Block until signal received
	fmt.Println("\nShutdown signal received. Stopping all goroutines...")
	stop.Store(true) // Signal all goroutines to exit their loops
	fmt.Println("Waiting for goroutines to finish their current iteration...")
	wg.Wait() // Wait for all to complete
	fmt.Println("All goroutines stopped. Clean exit.")
}

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
