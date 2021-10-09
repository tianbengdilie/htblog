package auth

import (
	"htblog/internal/models/usermod"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	NickName string `json:"nickname"`
}

func dbUser2clientUser(dbUser *usermod.UserInfo) UserInfo {
	return UserInfo{
		NickName: dbUser.NickName,
	}
}

func Init(router *gin.Engine) {
	router.POST("/api/v1/auth/login", Login)
	router.POST("/api/v1/auth/register", Register)
	router.POST("/api/v1/auth/logout", Logout)
	router.POST("/api/v1/auth/get_user_info", GetUserInfo)
}
