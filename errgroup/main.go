package errgroup

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

var URLs = []string{
	"http://www.example.org",
	"http://www.cnn.com",
	"https://www.iana.org/help/foobar", // Intentionally 404 to demonstrate error handling
}

// tryRunWithLimit demonstrates non-blocking work submission with errgroup.TryGo.
// Unlike g.Go() which blocks when the limit is reached, TryGo() returns false immediately,
// allowing the caller to handle rejections (retry, queue, drop, etc).
func tryRunWithLimit() {
	var g errgroup.Group
	g.SetLimit(2) // Maximum 2 goroutines running concurrently

	// Task 1: runs for 5 seconds - fills first slot
	g.Go(func() error {
		fmt.Println("Running Task 1")
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Task 1 done")
		}
		return nil
	})

	// Task 2: runs for 10 seconds - fills second slot
	// Both slots now occupied until Task 1 completes at t=5s
	g.Go(func() error {
		fmt.Println("Running Task 2")
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Task 2 done")
		}
		return nil
	})

	// Task 3: Define once, submit with TryGo in retry loop
	task3 := func() error {
		fmt.Println("Running Task 3")
		time.Sleep(2 * time.Second)
		fmt.Println("Task 3 done")
		return nil
	}

	// TryGo returns false when limit is reached (both slots occupied).
	// Loop retries every second until a slot opens (when Task 1 completes at t=5s).
	// This demonstrates non-blocking submission with explicit retry logic.
	for !g.TryGo(task3) {
		fmt.Println("Task 3 rejected - retrying in 1s...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Task 3 accepted")

	// Wait for all tasks to complete
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

// runWithLimit demonstrates concurrency limiting with errgroup.SetLimit.
// Spawns 50 goroutines but limits concurrent execution to 4 at a time.
// Uses atomic counter to track active goroutines and verify the limit is respected.
func runWithLimit() {
	var counter atomic.Int32 // Tracks number of currently active goroutines
	var id atomic.Int32      // Assigns unique ID to each goroutine
	var g errgroup.Group
	g.SetLimit(4) // Maximum 4 goroutines running concurrently

	// Start monitoring goroutine that prints active count every second
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// Counter should never exceed 4 due to SetLimit
				fmt.Printf("counter: %d\n", counter.Load())
			case <-done:
				return // Stop monitoring when all work completes
			}
		}
	}()

	// Submit 50 tasks - SetLimit ensures only 4 run simultaneously, rest queue
	for i := 0; i < 50; i++ {
		g.Go(func() error {
			// Assign unique ID and track start of work
			cur := id.Add(1) // Atomic increment, returns new value (1-50)
			fmt.Printf("goroutine %d\n", cur)

			counter.Add(1)                                        // Increment active count
			defer counter.Add(-1)                                 // Decrement when function exits
			time.Sleep(time.Duration(rand.Intn(2)) * time.Second) // Simulate work (0-1 seconds)
			return nil
		})
	}

	// Wait for all 50 goroutines to complete
	// SetLimit throttles execution, so this takes longer than without limit
	if err := g.Wait(); err != nil {
		fmt.Printf("error: %s", err)
	}

	close(done) // Stop the monitoring goroutine
}

// runWithContext demonstrates automatic context cancellation when one task fails.
// When errgroup.WithContext is used, the first goroutine to return an error triggers
// context cancellation, signaling other goroutines to stop. However, cancellation is
// cooperative - goroutines must check ctx.Done() to observe it. g.Wait() always waits
// for ALL goroutines to complete, even after cancellation.
func runWithContext() {
	type Result struct {
		statusCode int
	}
	resultChan := make(chan Result, 3)

	// WithContext creates a context that gets canceled when first error occurs
	g, ctx := errgroup.WithContext(context.Background())

	// Goroutine 1: simulates long-running work (5 seconds)
	g.Go(func() error {
		// Select races between completing work and observing cancellation.
		// When goroutine 3 errors at 1 second, ctx.Done() becomes ready and this
		// goroutine exits immediately without waiting the full 5 seconds.
		select {
		case <-time.After(5 * time.Second):
			// Work completed normally - send result
			resultChan <- Result{
				statusCode: http.StatusOK,
			}
		case <-ctx.Done():
			// Context canceled early - another goroutine errored
			err := ctx.Err()
			fmt.Printf("goroutine 1 - context canceled: %s\n", err)
			return err
		}
		return nil
	})

	// Goroutine 2: simulates long-running work (5 seconds)
	g.Go(func() error {
		// Select races between completing work and observing cancellation.
		// When goroutine 3 errors at 1 second, ctx.Done() becomes ready and this
		// goroutine exits immediately without waiting the full 5 seconds.
		select {
		case <-time.After(5 * time.Second):
			// Work completed normally - send result
			resultChan <- Result{
				statusCode: http.StatusOK,
			}
		case <-ctx.Done():
			// Context canceled early - another goroutine errored
			err := ctx.Err()
			fmt.Printf("goroutine 2 - context canceled: %s\n", err)
			return err
		}
		return nil
	})

	// Goroutine 3: errors early (1 second), triggering cancellation
	g.Go(func() error {
		select {
		case <-time.After(1 * time.Second):
			return fmt.Errorf("network error")
		case <-ctx.Done():
			// Context already canceled (unlikely in this scenario)
			err := ctx.Err()
			fmt.Printf("goroutine 3 - context canceled: %s\n", err)
			return err
		}
	})

	// Wait blocks until ALL three goroutines return, even though context is canceled.
	// Returns first non-nil error (from goroutine 3).
	// With context-aware select, goroutines 1 and 2 exit early (~1 second) when goroutine 3 errors.
	// This demonstrates the "fail fast" pattern - one error cancels outstanding work immediately.
	if err := g.Wait(); err != nil {
		fmt.Printf("run: error: %s\n", err)
	}

	// Close channel after all goroutines complete
	close(resultChan)

	// Collect any results that were sent before cancellation
	for result := range resultChan {
		fmt.Printf("result: %v\n", result)
	}
}

// fetchResults demonstrates collecting results from concurrent tasks with errgroup.
// Launches 3 concurrent HTTP fetches and gathers successful responses via a buffered channel.
// Even if one fetch fails, g.Wait() blocks until ALL goroutines complete, so we collect
// all successful results before processing errors.
func fetchResults() {
	// Result holds data from a successful HTTP fetch
	type Result struct {
		statusCode int
		partial    []byte
	}

	// Buffered channel prevents goroutines from blocking when sending results.
	// Size matches number of concurrent tasks - each successful fetch sends exactly once.
	resultChan := make(chan Result, len(URLs))
	g := &errgroup.Group{}

	// Launch concurrent fetches - each sends its result to the channel on success
	for _, url := range URLs {
		g.Go(func() error { // Go 1.22+: url is captured correctly per-iteration
			resp, err := http.DefaultClient.Get(url)
			if err != nil {
				return fmt.Errorf("error: url: %s, error: %w", url, err) // Network/connection error
			}
			defer resp.Body.Close() // Clean up resources
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error: url: %s, status: %d", url, resp.StatusCode) // HTTP error
			}
			d, _ := io.ReadAll(resp.Body)
			resultChan <- Result{ // Send result to channel - non-blocking due to buffer
				statusCode: resp.StatusCode,
				partial:    d[:256], // Read partial body to avoid large allocations
			}
			return nil // Success
		})
	}

	// Wait for ALL goroutines to complete (even if some return errors).
	// Returns first non-nil error encountered, but doesn't stop other goroutines.
	if err := g.Wait(); err != nil {
		fmt.Printf("error: %s\n", err)
	}

	// Close channel to signal no more results coming - allows range loop to terminate
	close(resultChan)

	// Collect all successful results from channel
	var results []Result
	for result := range resultChan {
		results = append(results, result)
	}

	fmt.Printf("collected %d results\n", len(results))
}

// fetchURLs demonstrates the fan-out pattern with errgroup.
// Launches 3 concurrent HTTP fetches. If any returns an error, g.Wait() returns
// that error. The errgroup automatically waits for all goroutines to complete.
func fetchURLs() {
	g := &errgroup.Group{} // Create errgroup to manage goroutines with error propagation

	// Launch concurrent fetches - each runs in its own goroutine
	for _, url := range URLs {
		g.Go(func() error { // Go 1.22+: url is captured correctly per-iteration
			resp, err := http.DefaultClient.Get(url)
			defer resp.Body.Close() // Clean up resources
			if err != nil {
				return err // Network/connection error
			}
			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("error: url: %s, status: %d", url, resp.StatusCode) // HTTP error
			}
			fmt.Printf("status: url: %s, %d\n", url, resp.StatusCode)
			return nil // Success
		})
	}

	// Wait for all fetches to complete; returns first error encountered
	if err := g.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
