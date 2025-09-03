package users

import (
	"net/http"
	"strconv"
	"trainnig-api-poc/api/users/entities"
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
	userId := c.Param("userId")
	c.JSON(200, gin.H{
		"message": "User " + userId + " updated",
	})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	c.JSON(200, gin.H{
		"message": "User " + userId + " deleted",
	})
}
