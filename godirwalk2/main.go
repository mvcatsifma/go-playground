package main

import (
	"errors"
	"fmt"
	"github.com/karrick/godirwalk"
	"log"
	"strings"
)

func main() {
	dirname := "/mnt/d/ws"
	result := &WalkResult{}
	err := godirwalk.Walk(dirname, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			log.Println(path)
			if de.IsSymlink() {
				log.Printf("skip symbolic link: %v\n", path)
				return godirwalk.SkipThis
			}
			if strings.Contains(path, "onion«"){
				return ErrFatal
			}
			result.Visited++
			return nil
		},
		// ErrFatal stops the entire walk; all other errors skip the offending node.
		ErrorCallback: func(path string, err error) godirwalk.ErrorAction {
			if errors.Is(ErrFatal, err) {
				return godirwalk.Halt
			}
			return godirwalk.SkipNode
		},
	})
	if errors.Is(ErrFatal, err) {
		log.Fatalln(err)
	}
	log.Println("ok")
}

var ErrFatal = fmt.Errorf("fatal error")

type WalkResult struct {
	Err error
	Visited int
}