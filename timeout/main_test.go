package main

import (
	"context"
	"testing"
	"time"
)

func Test_ShouldReturnError_WhenStatError(t *testing.T) {
	task := &Task{}
	fs := &BrokenFs{}

	ctx := context.Background()
	result := handleTask(task, fs, ctx)

	if result.pathErrors != 1 {
		t.Fatal("expected one path error")
	}
	if result.permissionErrors != 1 {
		t.Fatal("expected one permission error")
	}
}

func Test_ShouldReturnCanceled_WhenTimeout(t *testing.T) {
	task := &Task{}
	fs := &StubFs{}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	result := handleTask(task, fs, ctx)

	if !result.canceled {
		t.Fatal("expected canceled")
	}
}

func Test_ShouldReturnCanceled_WhenCanceled(t *testing.T) {
	task := &Task{}
	fs := &StubFs{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result := handleTask(task, fs, ctx)

	if !result.canceled {
		t.Fatal("expected canceled")
	}
}

