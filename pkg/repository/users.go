package repository

import (
	"ToDoListDbGormGin/db"
	"ToDoListDbGormGin/logger"
	"ToDoListDbGormGin/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error grtting all users: %s\n", err.Error())
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserById] error grtting user by id: %s\n", err.Error())
		return user, err
	}

	return user, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error

	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %s\n", err.Error())
		return err
	}

	return nil
}

func GetUserByUserName(username string) (user models.User, err error) {
	if err = db.GetDBConn().Where("username = ?", username).First(&user).Error; err != nil {
		logger.Error.Printf("[repository.GetUserByUserName] error grtting user by username: %s\n", err.Error())
		return user, err
	}

	return user, nil
}

func UpdateUserByID(id int, updatedUser models.User) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserByID] error grtting user by id: %s\n", err.Error())
		return user, err
	}

	err = db.GetDBConn().Model(&user).Updates(updatedUser).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUserByID] error updating user by id: %s\n", err.Error())
		return user, err
	}

	return user, nil
}

func DeleteUserById(id int) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUserById] error delete user by id: %s\n", err.Error())
		return user, err
	}

	err = db.GetDBConn().Delete(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUserById] error delete user by id: %s\n", err.Error())
		return user, err
	}

	return user, nil
}
