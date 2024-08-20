package main

import (
	"fmt"
	"log"
	"os"
	"taskTracker/internal/cli"
)

func main() {
	fmt.Println("Task Tracker")
	if len(os.Args) > 1 {
		cli.HandleCommand(os.Args)
	}
	log.Fatal("command not provided")
}
