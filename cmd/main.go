package main

import (
	"fmt"
	"log"
	"os"
	"taskTracker/internal/cli"
)

func main() {
	fmt.Println("Task Tracker")
	if len(os.Args) <= 1 {
		log.Fatal("command not provided")
	}
	cli.HandleCommand(os.Args)
}
