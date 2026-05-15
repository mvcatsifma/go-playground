package main

import (
	"embed"
	"fmt"
	"log"
)

//go:embed testdata
var f embed.FS

func main() {
	// Print directory entries
	dir, err := f.ReadDir("testdata")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range dir {
		log.Println(entry.Name())
	}

	// Print contents of file 1.txt
	content, err := f.ReadFile("testdata/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
