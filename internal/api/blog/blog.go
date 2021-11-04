package blog

import (
	"htblog/internal/models/blogmod"
	mhelper "htblog/internal/models/helper"
	"htblog/internal/pkg/tginutils"
	"htblog/internal/terror"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Blog struct {
	BlogID    string    `json:"blog_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Editor    string    `json:"editor"`
	ParentID  string    `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SaveBlogReq struct {
	BlogID   string `json:"blog_id"`
	Title    string `json:"title"`
	Type     string `json:"type"` // md | rich-text
	Content  string `json:"content"`
	ParentID string `json:"parent_id"`
	Editor   string `json:"editor"`
}

type GetBlogReq struct {
	BlogID string `json:"blog_id"`
}

func SaveBlog(c *gin.Context) {
	tlog := tginutils.GetLog(c)

	req := SaveBlogReq{}
	err := c.BindJSON(&req)
	if err != nil {
		tlog.Info("c.BindJSON err : ", err)
		c.JSON(http.StatusOK, terror.ErrInvalidParam)
		return
	}

	err = blogmod.SaveBlog(blogmod.Blog{
		BlogID:   req.BlogID,
		Title:    req.Title,
		Content:  req.Content,
		ParentID: req.ParentID,
		Editor:   req.Editor,
	})
	if err != nil {
		tlog.Error("blogmod.SaveBlog err : ", err)
		c.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}

	c.JSON(http.StatusOK, terror.Succeed)
}

func GetBlog(c *gin.Context) {
	tlog := tginutils.GetLog(c)

	req := GetBlogReq{}
	err := c.BindJSON(&req)
	if err != nil {
		tlog.Info("c.BindJSON err : ", err)
		c.JSON(http.StatusOK, terror.ErrInvalidParam)
		return
	}

	dbBlog, err := blogmod.GetBlog(req.BlogID)
	if mhelper.IsNotFound(err) {
		tlog.Warnf("req[%s] not found", req.BlogID)
		c.JSON(http.StatusOK, terror.ErrDataNotFound)
		return
	}
	if err != nil {
		tlog.Error("blogmod.GetBlog err : ", err)
		c.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}

	blog := Blog{
		BlogID:    dbBlog.BlogID,
		Title:     dbBlog.Title,
		Content:   dbBlog.Content,
		Editor:    dbBlog.Editor,
		ParentID:  dbBlog.ParentID,
		CreatedAt: dbBlog.CreatedAt,
		UpdatedAt: dbBlog.UpdatedAt,
	}

	c.JSON(http.StatusOK, terror.Msg{
		MsgBase: terror.Succeed,
		Data:    blog,
	})
}

type BlogList struct {
	Title     string `json:"title"`
	BlogID    string `json:"blog_id"`
	ChildBlog []*BlogList
}

func GetBlogList(c *gin.Context) {
	tlog := tginutils.GetLog(c)

	allBlogs, err := blogmod.GetAllBlogBrief()
	if err != nil {
		tlog.Error("blogmod.GetAllBlogs err : ", err)
		c.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}

	rootBlogList := make([]*BlogList, 0)
	builder := make(map[string]*BlogList)
	for _, brief := range allBlogs {
		if _, ok := builder[brief.BlogID]; !ok {
			builder[brief.BlogID] = &BlogList{
				BlogID:    brief.BlogID,
				ChildBlog: make([]*BlogList, 0),
			}
		}
		builder[brief.BlogID].Title = brief.Title

		if brief.ParentID != "" {
			if _, ok := builder[brief.ParentID]; !ok {
				builder[brief.ParentID] = &BlogList{
					BlogID:    brief.ParentID,
					ChildBlog: make([]*BlogList, 0),
				}
			}
			builder[brief.ParentID].ChildBlog = append(builder[brief.ParentID].ChildBlog, builder[brief.BlogID])
		} else {
			rootBlogList = append(rootBlogList, builder[brief.BlogID])
		}
	}

	c.JSON(http.StatusOK, terror.Msg{
		MsgBase: terror.Succeed,
		Data:    rootBlogList,
	})
}
