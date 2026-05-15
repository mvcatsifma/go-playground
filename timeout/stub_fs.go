package main

import (
	"github.com/spf13/afero/mem"
	"os"
	"time"
)

// StubFs is a stub implementation of the Fs interface
type StubFs struct {
}

func (m StubFs) Stat(name string) (os.FileInfo, error) {
	time.Sleep(1 * time.Second) // simulate work being done
	d := mem.CreateFile("test")
	i := &mem.FileInfo{FileData: d}
	return i, nil
}