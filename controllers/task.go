package task

import (
	"github.com/gentom/gincrud/models"
)

type Task struct {
}

func TASK() Task {
	return Task{}
}

func (c Task) AddTask(text string) {
	repo := task.TaskRepository()
	repo.Create(text)
}

func (c Task) GetAllTasks() interface{} {
	repo := task.TaskRepository()
	tasks := repo.GetALL()
	return tasks
}
