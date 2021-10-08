package auth

import "github.com/gin-gonic/gin"

func Init(router *gin.Engine) {
	router.POST("/api/v1/auth/login", Login)
	router.POST("/api/v1/auth/register", Register)
	router.POST("/api/v1/auth/logout", Logout)
}
