package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
	"taskTracker/internal/task"
)

func HandleCommand(c []string) {

	task := task.Tasks{}

	switch c[0] {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := cmd.String("description", "", "Task description")

		if len(os.Args) < 3 {
			log.Fatal("Please provide a task description")
		}

		result := task.NewTask(*description)

		fmt.Println(result)

	case "exit":
		os.Exit(0)

	default:
		log.Fatalf("Unknown command %v", c[0])

	}
}
