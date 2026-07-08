package main

import (
	"context"
	"errors"
	fspkg "io/fs"
	"log"
	"os"
	"os/signal"
	"time"

	"golang.org/x/sync/errgroup"
)

const limit = 2

// main demonstrates context-aware directory walking with concurrency limiting, OS signal handling,
// and timeout detection. Uses errgroup.SetLimit to bound concurrent tasks and distinguishes
// between user cancellation (SIGINT) and context timeout.
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Register signal handler to cancel context on Ctrl+C
	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)
	go func() {
		<-osInterruptChannel
		log.Println("signal received, shutting down")
		cancel()
	}()

	root := "/Users/tdelnoij/code/github.com/mvcatsifma/go-playground/timeout"
	fs := os.DirFS(root)

	tasks := []*Task{
		{id: 1, path: "testdata/level1/a"},
		{id: 2, path: "testdata/level1/b"},
		{id: 3, path: "testdata/level1/c"},
		{id: 4, path: "testdata/level1/restricted"},
	}

	runTasks(ctx, fs, tasks)
}

// runTasks processes tasks concurrently with context awareness and concurrency limiting.
// Returns after all tasks complete or context is canceled/times out.
func runTasks(ctx context.Context, fs fspkg.FS, tasks []*Task) {
	var g errgroup.Group
	g.SetLimit(limit) // Maximum 2 goroutines running concurrently

	// Launch all tasks concurrently (respecting SetLimit)
	for _, task := range tasks {
		g.Go(func() error {
			result := runTask(ctx, fs, task)
			log.Printf("task[%d] result: canceled=%v timeout=%v foundTarget=%v visited=%d err=%v\n",
				task.id, result.canceled, result.timeout, result.foundTarget, result.visited, result.err)
			return nil
		})
	}

	log.Println("all tasks started, waiting for completion or interrupt")
	_ = g.Wait()
	log.Println("all tasks complete, shutting down")
}

// runTask walks task.path recursively using fs.WalkDir, checking ctx.Done() on every entry
// to stop promptly on timeout or cancellation. Non-critical errors (permission, path) are logged
// but don't halt the walk. Returns TaskResult with visit count and cancellation/timeout status.
func runTask(ctx context.Context, fs fspkg.FS, task *Task) *TaskResult {
	log.Printf("task[%d]: handling now\n", task.id)

	result := &TaskResult{taskId: task.id}

	err := fspkg.WalkDir(fs, task.path, func(path string, d fspkg.DirEntry, walkErr error) error {
		// Handle errors from WalkDir (e.g., permission denied on directory read)
		if walkErr != nil {
			result.err = walkErr
			log.Printf("ERROR: task[%d]: %s\n", task.id, walkErr)
			return nil // Log error but continue walking
		}

		// Skip "restricted" directories proactively before WalkDir attempts to read them.
		// Demonstrates fs.SkipDir: skip this directory's contents but continue with siblings.
		if d.IsDir() && d.Name() == "restricted" {
			log.Printf("task[%d]: skipping restricted directory: %s\n", task.id, path)
			return fspkg.SkipDir
		}

		// Search for sentinel file by name. When found, stop entire walk immediately.
		// Demonstrates fs.SkipAll: terminate walk completely (successful early exit).
		if !d.IsDir() && d.Name() == "sentinel-deadbeef-marker.txt" {
			log.Printf("task[%d]: found sentinel file, stopping walk: %s\n", task.id, path)
			result.visited++
			result.foundTarget = true
			return fspkg.SkipAll
		}

		// Check context before processing each entry (timeout or cancellation detection)
		select {
		case <-ctx.Done():
			err := ctx.Err()
			if errors.Is(err, context.Canceled) {
				result.canceled = true
				log.Printf("task[%d]: canceled by user interrupt\n", task.id)
				return TaskCanceled
			}

			result.timeout = true
			log.Printf("task[%d]: timeout after 60 seconds\n", task.id)
			return TaskTimeout
		default:
			// Process entry: get info and log visit
			fileInfo, err := d.Info()
			if err != nil {
				result.err = err
				log.Printf("ERROR: task[%d]: %s\n", task.id, err)
				return nil
			}

			log.Printf("task[%d] visited: %s %s\n", task.id, fileInfo.Mode(), path)
			result.visited++
			return nil
		}
	})

	if err != nil {
		result.err = err
	}

	return result
}
