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

	root := "/Users/tdelnoij/code/github.com/mvcatsifma/go-playground/timeout"
	fs := os.DirFS(root)

	var g errgroup.Group
	g.SetLimit(2) // Maximum 2 goroutines running concurrently

	// Register signal handler to cancel context on Ctrl+C
	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)
	go func() {
	readInterruptLoop:
		for {
			select {
			case <-osInterruptChannel:
				log.Println("signal received, shutting down")
				cancel()
				break readInterruptLoop
			}
		}
	}()

	// Define tasks to process concurrently (max 2 at a time due to SetLimit)
	tasks := []*Task{
		{id: 1, path: "testdata/level1/a"},
		{id: 2, path: "testdata/level1/b"},
		{id: 3, path: "testdata/level1/c"},
	}
	for _, task := range tasks {
		g.Go(func() error {
			result := handleTask(task, fs, ctx)
			log.Printf("task[%d] result: canceled=%v timeout=%v visited=%d err=%s\n",
				task.id, result.canceled, result.timeout, result.visited, result.err)
			return nil
		})
	}

	log.Println("all tasks started, waiting for completion or interrupt")

	_ = g.Wait()

	log.Println("all tasks complete, shutting down")
}

// handleTask walks task.path recursively using fs.WalkDir, checking ctx.Done() on every entry
// to stop promptly on timeout or cancellation. Non-critical errors (permission, path) are logged
// but don't halt the walk. Returns TaskResult with visit count and cancellation/timeout status.
func handleTask(task *Task, fs fspkg.FS, ctx context.Context) *TaskResult {
	log.Printf("task[%d]: handling now\n", task.id)

	result := &TaskResult{taskId: task.id}

	_ = fspkg.WalkDir(fs, task.path, func(path string, d fspkg.DirEntry, err error) error {
		if err != nil {
			if errors.Is(err, TaskCanceled) || errors.Is(err, TaskTimeout) {
				result.err = err.Error()
				return err // Halt walk on cancellation or timeout
			}
			// Log other errors but continue walking (keeps first error only)
			log.Printf("ERROR: task[%d]: %s\n", task.id, err)
			if result.err == "" {
				result.err = err.Error()
			}
			return nil // Continue despite error
		}

		// Check context before processing each entry
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
				return err
			}
			log.Printf("task[%d] visited: %s %s\n", task.id, fileInfo.Mode(), path)
			result.visited++
			return nil
		}
	})

	log.Printf("task[%d]: handling complete\n", task.id)

	return result
}
