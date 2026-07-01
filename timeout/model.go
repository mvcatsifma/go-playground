package main

import (
	"errors"
)

type Task struct {
	id   int64
	path string
}

type TaskResult struct {
	canceled bool
	timeout  bool
	taskId   int64
	visited  int64
	err      string
}

var TaskCanceled = errors.New("task canceled")

var TaskTimeout = errors.New("task timeout")
