package main

import (
	"errors"
	"os"
)

type Token struct {
}

type Task struct {
	id   int64
	path string
}

type TaskResult struct {
	canceled   bool
	pathErrors int64
	permissionErrors int64
	//elasticErrors int64 todo
	taskId     int64
	visited    int64
}

var TaskCanceled = errors.New("task canceled")

type Fs interface {

	// Stat returns a FileInfo describing the named file.
	// If there is an error, it will be of type *PathError.
	Stat(name string) (os.FileInfo, error)
}

