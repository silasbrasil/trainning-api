package users

import (
	"net/http"
	"strconv"
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

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "User created",
	})
}

func GetUserById(c *gin.Context) {
	userIdParam := c.Param("userId")
	userId, _ := strconv.ParseUint(userIdParam, 10, 64)

	user := usecases.GetUserById(userId)

	c.JSON(http.StatusOK, user)
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
