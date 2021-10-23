package blog

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	g := r.Group("/api/v1/blog")
	g.POST("/save_blog", SaveBlog)
	g.GET("/get_blog", GetBlog)
	g.POST("/save_img", SaveImg)
	g.GET("/get_img", GetImg)
}
