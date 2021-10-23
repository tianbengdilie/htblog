package blog

import (
	"htblog/internal/models/blogmod"
	"htblog/internal/pkg/tginutils"
	"time"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	BlogID    string    `json:"blog_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Editor    string    `json:"editor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SaveBlogReq struct {
	BlogID  string `json:"blog_id"`
	Title   string `json:"title"`
	Type    string `json:"type"` // md | rich-text
	Content string `json:"content"`
	Editor  string `json:"editor"`
}

type GetBlogReq struct {
	BlogID string `json:"blog_id"`
}

func SaveBlog(c *gin.Context) {
	tlog := tginutils.GetLog(c)

	req := SaveBlogReq{}
	err := c.BindJSON(&req)
	if err != nil {
		tlog.Error("c.BindJSON")
	}

	err := blogmod.SaveBlog(blogmod.Blog{
		BlogID: ,
	})
}

func GetBlog(c *gin.Context) {

}
