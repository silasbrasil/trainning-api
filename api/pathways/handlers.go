package pathways

import (
	"github.com/gin-gonic/gin"
)

func ListPathways(c *gin.Context) {
	c.JSON(200, gin.H{
		"pathways": []string{"pathway1", "pathway2"},
	})
}

func GetPathwayById(c *gin.Context) {
	pathwayId := c.Param("pathwayId")
	c.JSON(200, gin.H{
		"pathway": pathwayId,
	})
}
