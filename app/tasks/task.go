package tasks

import (
	"github.com/gofrs/uuid"
)

type StatusTask string
const (
	ToDo  StatusTask = "TODO"
	Doing StatusTask = "DOING"
	Done  StatusTask = "DONE"
)

func (s *StatusTask) String() string {
	return string(*s)
}

func NewStatus(status string) StatusTask {
	switch status {
	case string(ToDo):
		return ToDo
	case string(Doing):
		return Doing
	case string(Done):
		return Done
	default:
		return ""
	}
}

// Task Domain struct for manage tasks
type Task struct {
	Id          string
	Title       string
	Description string
	Status      StatusTask
}

func NewTask(title string, description string, status StatusTask) *Task {
	id, _ := uuid.NewV4()
	return &Task{
		Id:          id.String(),
		Title:       title,
		Description: description,
		Status:      status,
	}
}
