package task

import (
	"math"
	"math/rand"
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

}

func (t *Tasks) Save(filename string) error {

}
