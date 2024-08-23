package cli

import (
	"flag"
	"log"
	"taskTracker/internal/task"
	"taskTracker/util"
)

const filename = "data/tasks.json"

func HandleCommand(commands []string) {
	args := commands[2:]
	task := task.Tasks{}

	err := task.Load(filename)
	util.LogError(err)

	switch commands[1] {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := cmd.String("description", "", "Task description")

		cmd.Parse(args)

		if *description == "" {
			log.Printf("Please provide a task description, have: %s", *description)
			break
		}

		task.NewTask(*description)

		log.Println("Task added succesfully")

	case "update":
		cmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := cmd.Int("id", 0, "id of the task")
		status := cmd.String("status", "pending", "status of the task")

		cmd.Parse(args)

		err := task.Update(*id, *status)
		util.LogError(err)

	case "delete":
		cmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := cmd.Int("id", 0, "id of the task")

		cmd.Parse(args)

		err := task.Delete(*id)
		util.LogError(err)

	case "print":

		cmd := flag.NewFlagSet("filter", flag.ExitOnError)
		filter := cmd.String("filter", "", "filter tasks")

		cmd.Parse(args)

		if *filter == "" {
			log.Println("Filter is empty")
			break
		}

		task.Print(*filter)

	default:
		log.Fatalf("Unknown command %v", commands)

	}
	err = task.Save(filename)
	util.LogError(err)

}
