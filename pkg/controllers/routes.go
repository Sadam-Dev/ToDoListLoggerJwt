package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunRoutes() error {
	router := gin.Default()

	router.GET("/ping", PinPong)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp)
		auth.POST("/sign-in", SignIn)
	}

	userG := router.Group("/users")
	{
		userG.GET("", GetAllUsers)
		userG.GET("/:id", GetUserById)
		userG.POST("", CreateUser)
		userG.PUT("/:id", UpdateUserByID)
		userG.DELETE("/:id", DeleteUserById)
	}

	tasksG := router.Group("/tasks", checkUserAuthentication)
	{
		tasksG.GET("", GetAllTasks)
		tasksG.GET("/:id", GetTaskByID)
		tasksG.POST("", CreateTask)
		tasksG.PUT("/:id", UpdateTaskByID)
		tasksG.DELETE("/:id", DeleteTaskByID)
	}

	err := router.Run(":5555")
	if err != nil {
		return err
	}

	return nil
}

func PinPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
