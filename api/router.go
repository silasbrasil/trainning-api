package api

import (
	"trainnig-api-poc/api/pathways"
	"trainnig-api-poc/api/pathways/participants"
	"trainnig-api-poc/api/users"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "running",
		})
	})

	users.SetRoute(r)
	pathways.SetRoute(r)
	participants.SetRoute(r)

	return r
}
