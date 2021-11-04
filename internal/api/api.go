package api

import (
	"htblog/internal/api/auth"
	"htblog/internal/api/blog"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	auth.Init(r)
	blog.Init(r)
}
