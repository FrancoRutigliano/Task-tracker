package cli

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"taskTracker/internal/task"
)

const filename = "data/tasks.json"

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	task := task.Tasks{}

	if err := task.Load(filename); err != nil {
		log.Fatal("error loading tasks: ", err)
	}

	for {
		fmt.Print("> ")
		scanner.Scan()

		input := scanner.Text()
		if input == "" {
			continue
		}

		args := strings.Split(input, " ")
		command := args[0]
		HandleCommand(task, command, args[1:])
	}

}

func HandleCommand(task task.Tasks, command string, args []string) {

	switch command {
	case "add":
		cmd := flag.NewFlagSet("add", flag.ExitOnError)
		description := cmd.String("description", "", "Task description")

		cmd.Parse(args)

		if *description == "" {
			log.Printf("Please provide a task description, have: %s", *description)
			break
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
		log.Print("Unknown command %v", command)

	}

}
