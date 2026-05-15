package main

import (
	"log"
)

// See log.SetFlags
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Llongfile | log.Lmsgprefix)
	log.SetPrefix("crawler-0: ")
	log.Println("Hello world!")
}
