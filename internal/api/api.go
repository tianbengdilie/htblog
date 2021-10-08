package api

import (
	"htblog/internal/api/auth"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	auth.Init(r)
}
