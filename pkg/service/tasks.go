package service

import (
	"ToDoListDbGormGin/models"
	"ToDoListDbGormGin/pkg/repository"
)

func GetAllTasks() (tasks []models.Task, err error) {
	tasks, err = repository.GetAllTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTaskByID(id int) (task models.Task, err error) {
	task, err = repository.GetTasksByID(id)
	if err != nil {
		return task, err
	}

	return task, nil
}

func CreateTask(task models.Task) error {
	err := repository.CreateTask(task)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTaskByID(id int, updatedTask models.Task) (task models.Task, err error) {
	task, err = repository.UpdateTaskByID(id, updatedTask)
	if err != nil {
		return task, err
	}

	return task, nil
}

func DeleteTaskByID(id int) (task models.Task, err error) {
	task, err = repository.DeleteTaskByID(id)
	if err != nil {
		return task, err
	}

	return task, nil
}
