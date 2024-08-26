package service

import (
	"ToDoListDbGormGin/models"
	"ToDoListDbGormGin/pkg/repository"
	"ToDoListDbGormGin/utils"
	"errors"
	"gorm.io/gorm"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id int) (user models.User, err error) {
	user, err = repository.GetUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	_, err = repository.GetUserByUserName(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	user.Password = utils.GenerateHash(user.Password)

	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserByID(id int, updatedUser models.User) (user models.User, err error) {
	user, err = repository.UpdateUserByID(id, updatedUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUserById(id int) (user models.User, err error) {
	user, err = repository.DeleteUserById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}
