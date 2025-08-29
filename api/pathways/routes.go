package pathways

import "github.com/gin-gonic/gin"

func SetRoute(r *gin.Engine) {
	r.GET("/pathways", ListPathways)
	r.GET("/pathways/:pathwayId", GetPathwayById)
}
