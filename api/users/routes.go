package users

import "github.com/gin-gonic/gin"

func SetRoute(r *gin.Engine) {
	r.GET("/users", ListUsers)
	r.POST("/users", CreateUser)
	r.GET("/users/:userId", GetUserById)
	r.PATCH("/users/:userId", UpdateUser)
	r.DELETE("/users/:userId", DeleteUser)
}
