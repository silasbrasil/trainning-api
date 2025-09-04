package users

import (
	"fmt"
	"net/http"
	"strconv"
	"trainnig-api-poc/api/users/entities"
	"trainnig-api-poc/api/users/repositories"
	"trainnig-api-poc/api/users/usecases"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data any  `json:"data"`
	Meta Meta `json:"meta,omitempty"`
}

type Meta struct {
	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

func ListUsers(c *gin.Context) {
	users := usecases.ListUsers()

	response := Response{
		Data: users,
	}

	c.JSON(http.StatusOK, response)
}

func GetUserById(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, _ := strconv.ParseUint(userIdParam, 10, 64)

	user := usecases.GetUserById(userId)

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user entities.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := usecases.CreateUser(user.Name, user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func UpdateUser(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, _ := strconv.ParseUint(userIdParam, 10, 64)
	updatedPayload := entities.User{}

	if err := c.ShouldBindJSON(&updatedPayload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Payload: %+v\n", updatedPayload)

	userRepo := repositories.NewUserRepository()
	existingUser := userRepo.GetById(userId)

	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	existingUser.Name = updatedPayload.Name
	existingUser.Email = updatedPayload.Email

	_, err := userRepo.Update(userId, existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, existingUser)
}

func DeleteUser(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, _ := strconv.ParseUint(userIdParam, 10, 64)

	userRepo := repositories.NewUserRepository()
	existingUser := userRepo.GetById(userId)

	if existingUser == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := userRepo.Delete(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.Status(http.StatusNoContent)
}
