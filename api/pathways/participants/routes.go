package participants

import "github.com/gin-gonic/gin"

func SetRoute(r *gin.Engine) {
	r.GET("/pathways/:pathwayId/participants", ListParticipants)
	r.GET("/pathways/:pathwayId/participants/:participantId", GetParticipantById)
}
