package task

import (
	"encoding/json"
	"errors"
	"fmt"
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
		Id:          rand.Intn(math.MaxInt32),
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
