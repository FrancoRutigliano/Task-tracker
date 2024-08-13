package cli

import (
	"flag"
	"log"
	"os"
	"taskTracker/internal/task"
)

const filename = "data/tasks.json"

func HandleCommand(c []string) {

	task := task.Tasks{}

	if err := task.Load(filename); err != nil {
		log.Fatal("error loading tasks: ", err)
	}

	switch c[0] {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := cmd.String("description", "", "Task description")

		if len(os.Args) < 3 {
			log.Fatal("Please provide a task description")
		}

		task.NewTask(*description)

	case "exit":
		os.Exit(0)

	default:
		log.Fatalf("Unknown command %v", c[0])

	}
}
