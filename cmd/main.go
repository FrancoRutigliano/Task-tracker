package main

import (
	"log"
	"os"
	"taskTracker/internal/cli"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("Please try again, you have to provide two arguments, and have %v", len(os.Args))
	}
	cli.HandleCommand(os.Args[1:])
}
