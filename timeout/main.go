package main

import (
	"context"
	"dsen.nl/go-playground/timeout/system"
	"fmt"
	"github.com/karrick/godirwalk"
	"log"
	"os"
	"os/signal"
	"time"
)

const limit = 2

// todo: should be able to detect difference between
// process cancellation and timeout.
func main() {
	fs := system.NewAferoFs()
	sem := make(chan Token, limit)
	tasks := make(chan *Task, 1)
	//tasks <- &Task{id: 1, path: "D:\\"}
	tasks <- &Task{id: 1, path: "/mnt/d"}

	osInterruptChannel := make(chan os.Signal, 1)
	signal.Notify(osInterruptChannel, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	go func() {
		for task := range tasks {
			sem <- Token{}
			go func(task *Task) {
				result := handleTask(task, fs, ctx)
				fmt.Printf("task result: %+v\n", result)
				<-sem
			}(task)
		}
	}()

	// This read will block untill an OS signal (such as Ctrl-C) is received:
	go func() {
	readInterruptLoop:
		for {
			select {
			case sig := <-osInterruptChannel:
				switch sig {
				case os.Interrupt: // SIGINT received
					log.Println("SIGINT received, shutting down")
					close(tasks)
					cancel()
					break readInterruptLoop
				default:
					log.Printf("unknown signal received: %v", sig)
				}
			}
		}
	}()

	fmt.Println("all tasks started, wait for done or interrupt")
	for n := limit; n > 0; n-- {
		sem <- Token{}
	}

	fmt.Println("done")
}

func handleTask(task *Task, fs Fs, ctx context.Context) *TaskResult {
	log.Printf("task[%v]: handling now \n", task.id)

	result := &TaskResult{taskId: task.id}

	err := godirwalk.Walk(task.path, &godirwalk.Options{
		Callback: func(path string, entry *godirwalk.Dirent) error {
			select {
			case <-ctx.Done():
				return TaskCanceled
			default:
				fileInfo, err := fs.Stat(path)
				if err != nil {
					return err
				}

				log.Printf("%s %s\n", fileInfo.Mode(), path)
				result.visited++
				return nil
			}
		},
		ErrorCallback: func(s string, err error) godirwalk.ErrorAction {
			if err == TaskCanceled {
				result.canceled = true
				return godirwalk.Halt
			}
			if _, ok := err.(*os.PathError); ok {
				if os.IsPermission(err) {
					result.permissionErrors++
				}
				result.pathErrors++
			}
			// if elastic error... todo
			log.Printf("ERROR: task[%v]: %s\n", task.id, err)
			return godirwalk.SkipNode
		},
		Unsorted: true,
	})
	if err != nil {
		log.Printf("ERROR: task[%v]: %v\n", task.id, err)
	}

	log.Printf("task[%v]: handling complete \n", task.id)

	return result
}
