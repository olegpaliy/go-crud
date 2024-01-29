package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goleh.com/m/controllers"
	"goleh.com/m/initializres"
	"goleh.com/m/services"
)

func init() {
	initializres.LoadEnvVariables()
}

func main() {
	// Create a new instance of InMemoryUsersService
	userService := services.NewInMemoryUserService()

	// Create a new instance of UsersController
	userController := controllers.NewUserController(userService)

	// Create a new Gin router
	router := gin.Default()
	router.Use(cors.Default())

	// Define CRUD endpoints
	router.POST("/create", userController.CreateUser)
	router.GET("/getAll", userController.GetAllUsers)
	router.GET("/getById", userController.GetUserByID)
	router.PUT("/update", userController.UpdateUser)
	router.DELETE("/delete", userController.DeleteUser)

	// Start the server
	router.Run()
}
