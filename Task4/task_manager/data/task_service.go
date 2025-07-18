package data

import (
	"task_manager/models"
	"errors"

)

type TaskService struct {
	Tasks map[int]models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		Tasks: make(map[int]models.Task),
	}
}

func (ts *TaskService) AddTask(task models.Task) error {
	if _, exists := ts.Tasks[task.ID]; exists {
		return errors.New("task with this ID already exists")
	}
	ts.Tasks[task.ID] = task
	return nil
}

func (ts *TaskService) GetTask(id int) (models.Task, error) {
	if task, exists := ts.Tasks[id]; exists {
		return task, nil
	}
	return models.Task{}, errors.New("task not found")
}

func (ts *TaskService) UpdateTask(id int, updatedTask models.Task) error {
	if _, exists := ts.Tasks[id]; !exists {
		return errors.New("task not found")
	}
	ts.Tasks[id] = updatedTask
	return nil
}

func (ts *TaskService) DeleteTask(id int) error {
	if _, exists := ts.Tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(ts.Tasks, id)
	return nil
}
func (ts *TaskService) ListTasks() []models.Task {
	tasks := make([]models.Task, 0, len(ts.Tasks))
	for _, task := range ts.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
