package cli

import (
	"flag"
	"fmt"
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

		cmd.Parse(c[1:])

		if *description == "" || len(os.Args) < 3 {
			log.Fatalf("Please provide a task description, have: %s", *description)
		}

		task.NewTask(*description)
		if err := task.Save(filename); err != nil {
			log.Fatal("error saving the newTask: ", err)
		}

		log.Println("Task added succesfully")

	case "update":
		cmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := cmd.Int("id", 0, "id of the task")
		status := cmd.String("status", "pending", "status of the task")

		cmd.Parse(c[1:])

		task.Update(*id, *status)
		if err := task.Save(filename); err != nil {
			log.Fatal("error saving the new task: ", err)
		}

	case "print":

		cmd := flag.NewFlagSet("filter", flag.ExitOnError)
		filter := cmd.String("filter", "", "filter tasks")

		cmd.Parse(c[1:])

		task.Print(*filter)

	case "exit":
		fmt.Println("goodbye...")
		os.Exit(0)

	default:
		log.Fatalf("Unknown command %v", c[0])

	}

}
