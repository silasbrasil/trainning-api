package users

import "github.com/gin-gonic/gin"

func ListUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": []string{"user1", "user2"},
	})
}

func CreateUser(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "User created",
	})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"user": id,
	})
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
