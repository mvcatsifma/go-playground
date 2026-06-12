package errgroup

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

var URLs = []string{
	"http://www.example.org",
	"http://www.cnn.com",
	"https://www.iana.org/help/foobar", // Intentionally 404 to demonstrate error handling
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
