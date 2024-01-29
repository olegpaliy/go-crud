package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goleh.com/m/services"
)

type UsersController struct {
	usersService services.UsersService
}

func NewUserController(usersService services.UsersService) *UsersController {
	return &UsersController{
		usersService: usersService,
	}
}

func (c *UsersController) CreateUser(ctx *gin.Context) {
	var newUser services.User
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.usersService.CreateUser(newUser)
	ctx.Status(http.StatusCreated)
}

func (c *UsersController) GetAllUsers(ctx *gin.Context) {
	users := c.usersService.GetAllUsers()
	ctx.JSON(http.StatusOK, users)
}

func (c *UsersController) GetUserByID(ctx *gin.Context) {
	id := ctx.Query("id")
	user, err := c.usersService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UsersController) UpdateUser(ctx *gin.Context) {
	var updatedUser services.User
	if err := ctx.ShouldBindJSON(&updatedUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.usersService.UpdateUser(updatedUser)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UsersController) DeleteUser(ctx *gin.Context) {
	id := ctx.Query("id")
	err := c.usersService.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}
