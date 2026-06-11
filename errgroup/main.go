package errgroup

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

// fetchURLs demonstrates the fan-out pattern with errgroup.
// Launches 3 concurrent HTTP fetches. If any returns an error, g.Wait() returns
// that error. The errgroup automatically waits for all goroutines to complete.
func fetchURLs() {
	URLs := []string{
		"http://www.example.org",
		"http://www.cnn.com",
		"https://www.iana.org/help/foobar", // Intentionally 404 to demonstrate error handling
	}

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
