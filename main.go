package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("go-importd: ")

	config := parseFlags(os.Args[1:])

	serve(config)
}
