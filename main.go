package main

import (
	"ToDoListDbGormGin/db"
	"ToDoListDbGormGin/logger"
	"ToDoListDbGormGin/pkg/controllers"
)

func main() {
	logger.Init()

	err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}

	err = db.Migrate()
	if err != nil {
		panic(err)
	}

	err = controllers.RunRoutes()
	if err != nil {
		panic(err)
	}

}
