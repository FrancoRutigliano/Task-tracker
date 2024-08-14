package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const (
	pending    = "pending"
	complete   = "complete"
	inProgress = "in progress"
)

type Tasks []Task

func (t *Tasks) NewTask(description string) {
	task := Task{
		Id:          rand.Intn(math.MaxInt8),
		Description: description,
		Status:      pending,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	*t = append(*t, task)
}

func (t *Tasks) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			os.Create("data/tasks.json")
			os.WriteFile(filename, []byte("[]"), 0644)
			fmt.Println("archive created")
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("eof file")
	}

	if err = json.Unmarshal(file, t); err != nil {
		return err
	}

	return nil
}

func (t *Tasks) Save(filename string) error {
	data, err := json.MarshalIndent(t, "", "")
	if err != nil {
		return err
	}

	os.WriteFile(filename, data, 0644)

	return nil
}

func (t *Tasks) Update(id int, status string) {
	for i := range *t {
		if (*t)[i].Id == id && (*t)[i].Status != status {
			oldStatus := (*t)[i].Status
			(*t)[i].Status = status
			log.Printf("Update task %d - Status %s to %s", (*t)[i].Id, oldStatus, (*t)[i].Status)
			return
		}
		log.Fatal("error to update the task")
	}
}

func (t *Tasks) Print(state string) {
	var filterTask Tasks

	if len(*t) == 0 {
		fmt.Println("no tasks")
	}

	switch state {
	case "":
		filterTask = append(filterTask, *t...)
	case inProgress:
		for _, v := range *t {
			if v.Status == inProgress {
				filterTask = append(filterTask, v)
			}
		}
	case pending:
		for _, v := range *t {
			if v.Status == pending {
				filterTask = append(filterTask, v)
			}
		}
	case complete:
		for _, v := range *t {
			if v.Status == complete {
				filterTask = append(filterTask, v)
			}
		}

	default:
		log.Println("Unknown filter")
		return
	}

	fmt.Println(filterTask)
}
