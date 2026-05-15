package main

import (
	"os"
)

// BrokenFs is a implementation of the Fs interface
// that always returns an error.
type BrokenFs struct {
}

func (b BrokenFs) Stat(name string) (os.FileInfo, error) {
	return nil, &os.PathError{Op: "stat", Path: name, Err: os.ErrPermission}
}
