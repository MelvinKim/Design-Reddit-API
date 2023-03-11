package controller

import (
	"net/http"
	"strconv"

	usecase "github.com/MelvinKim/Design-Reddit-API/Usecase"
	"github.com/MelvinKim/Design-Reddit-API/entity"
	"github.com/gin-gonic/gin"
)

// UserController is responsible for handling the user input and output of the application
type UserController struct {
	userService usecase.UserService
}

// NewUserController creates a new UserController instance
func NewUserController(userService usecase.UserService) *UserController {
	return &UserController{userService: userService}
}

// CreateUser creates a new user
func (c *UserController) CreateUser(ctx *gin.Context) {
	var input entity.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.CreateUser(input.FirstName, input.LastName, input.Email, input.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data saved successfully": user})
}

// GetUser gets a user using the UserID
func (c *UserController) GetUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User ID"})
		return
	}

	user, err := c.userService.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

// ListUsers fetches all application users
func (c *UserController) ListUsers(ctx *gin.Context) {
	users, err := c.userService.ListUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"users": users})
}
