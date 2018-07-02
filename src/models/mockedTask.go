package models

import (
	"github.com/satori/go.uuid"
)

//MockedCreateTask is mocked CreateTask func
func MockedCreateTask(task Task, err error) {
	CreateTask = func(task Task) (Task, error) {
		return task, err
	}
}

//MockedGetTask is mocked GetTask func
func MockedGetTask(task Task, err error) {
	GetTask = func(id uuid.UUID) (Task, error) {
		return task, err
	}
}

//MockedDeleteTask is mocked DeleteTask func
func MockedDeleteTask(id uuid.UUID, err error) {
	DeleteTask = func(id uuid.UUID) error {
		return err
	}
}

//MockedGetTasks is mocked GetTasks func
func MockedGetTasks(task []Task, err error) {
	GetTasks = func() ([]Task, error) {
		return task, err
	}
}