package main

import (
	"errors"
)

type Task struct {
	id   int64
	path string
}

type TaskResult struct {
	canceled    bool
	timeout     bool
	foundTarget bool  // Set to true when sentinel file is found and walk stops early
	taskId      int64
	visited     int64
	err         error
}

var TaskCanceled = errors.New("task canceled")

var TaskTimeout = errors.New("task timeout")
