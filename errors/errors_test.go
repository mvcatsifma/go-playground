package main

import (
	"errors"
	"fmt"
	"log"
	"testing"
)

var XError = errors.New("xerror")

type NError struct {
	Err     error
	Message string
}

func (n *NError) Error() string {
	return fmt.Sprintf("nerror: %v", n.Err)
}

func (n *NError) Unwrap() error {
	return n.Err
}

func (n *NError) Temporary() bool {
	return true
}

func Test_Sentinel(t *testing.T) {
	f := func() error {
		return XError
	}
	err := f()
	if !errors.Is(err, XError) {
		t.Fatal("expected XError")
	}
}

func Test_SentinelWrapped(t *testing.T) {
	f := func() error {
		return fmt.Errorf("%w", XError)
	}
	err := f()
	if !errors.Is(err, XError) {
		t.Fatal("expected XError")
	}
}

func Test_WrappedError(t *testing.T) {
	f := func() error {
		return &NError{
			Err:     XError,
			Message: "something bad happened",
		}
	}
	err := f()
	if !errors.Is(err, XError) {
		t.Fatal("expected XError")
	}
	target := &NError{}
	if !errors.As(err, &target) {
		t.Fatal("expected NError")
	} else {
		if !target.Temporary() {
			t.Fatal("expected true")
		}
	}
}

func Test_Stacktrace(t *testing.T) {
	err := errors.New("oops")
	err2 := fmt.Errorf("error: %w", err)
	log.Println(err2)
}
