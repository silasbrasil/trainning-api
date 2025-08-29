package participants

import "github.com/gin-gonic/gin"

func ListParticipants(c *gin.Context) {
	c.JSON(200, gin.H{
		"participants": []string{"participant1", "participant2"},
	})
}

func GetParticipantById(c *gin.Context) {
	participantId := c.Param("participantId")
	c.JSON(200, gin.H{
		"participant": participantId,
	})
}
