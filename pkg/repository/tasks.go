package repository

import (
	"ToDoListDbGormGin/db"
	"ToDoListDbGormGin/logger"
	"ToDoListDbGormGin/models"
)

func GetAllTasks() (tasks []models.Task, err error) {
	err = db.GetDBConn().Find(&tasks).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllTasks] Error getting all tasks: %v", err)
		return nil, err
	}
	return tasks, nil
}

func GetTasksByID(id int) (task models.Task, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&task).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllTasks] Error getting task by id: %v", err)
		return task, err
	}

	return task, nil
}

func CreateTask(task models.Task) (err error) {
	if err = db.GetDBConn().Create(&task).Error; err != nil {
		logger.Error.Printf("[repository.CreateTask] Error creating task: %v", err)
		return err
	}

	return nil
}

func UpdateTaskByID(id int, updatedTask models.Task) (task models.Task, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&task).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateTaskByID] Error getting task by id: %v", err)
		return task, err
	}

	err = db.GetDBConn().Model(&task).Updates(updatedTask).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateTaskByID] Error updating task: %v", err)
		return task, err
	}

	return task, nil
}

func DeleteTaskByID(id int) (deleteTask models.Task, err error) {
	err = db.GetDBConn().Where("id = ?", id).Delete(&deleteTask).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteTaskByID] Error deleting task bt id: %v", err)
		return deleteTask, err
	}

	return deleteTask, nil
}
