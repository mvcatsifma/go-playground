package main

import (
	_ "embed"
	"log"
	"testing"
)


//go:embed testdata/1.txt
var content string

func TestName(t *testing.T) {
	log.Println(content)
}

