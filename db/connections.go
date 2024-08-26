package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

func ConnectToDB() error {
	connStr := "host=localhost port=5432 user=postgres dbname=todolist password=12345678"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return err
	}

	fmt.Println("Connected to DataBase")

	dbConn = db
	return nil
}

func GetDBConn() *gorm.DB {
	return dbConn
}
