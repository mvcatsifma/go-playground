package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	filesystem := &StdFS{}
	data, err := filesystem.ReadFile("./testdata/test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%s\n", string(data))
}

// StdFS implements fs.ReadFileFS and fs.StatFS using the real OS filesystem.
type StdFS struct {
}

func (m *StdFS) Open(name string) (fs.File, error) {
	return os.Open(name)
}

func (m *StdFS) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (m *StdFS) Stat(name string) (fs.FileInfo, error) {
	return os.Stat(name)
}
