package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
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
		err := task.Save(filename)
		util.LogError(err)

		log.Println("Task added succesfully")

	case "update":
		cmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := cmd.Int("id", 0, "id of the task")
		status := cmd.String("status", "pending", "status of the task")

		cmd.Parse(args)

		task.Update(*id, *status)
		if err := task.Save(filename); err != nil {
			log.Print("error saving the new task: ", err)
		}

	case "print":

		cmd := flag.NewFlagSet("filter", flag.ExitOnError)
		filter := cmd.String("filter", "", "filter tasks")

		cmd.Parse(args)

		task.Print(*filter)

	case "exit":
		fmt.Println("goodbye...")
		os.Exit(0)

	default:
		log.Printf("Unknown command %v", commands)

	}

}
