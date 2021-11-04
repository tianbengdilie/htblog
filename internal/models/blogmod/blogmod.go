package blogmod

import (
	"fmt"
	"htblog/internal/config"
	mhelper "htblog/internal/models/helper"
	"htblog/internal/pkg/utils"
)

const tableBlog = "t_blogs"

type Blog struct {
	mhelper.Model
	BlogID   string `gorm:"column:blog_id"`
	Title    string `gorm:"column:title"`
	Type     string `gorm:"column:type"`
	Content  string `gorm:"column:content"`
	ParentID string `gorm:"column:parent_id"` // 0表示根节点
	Editor   string `gorm:"column:editor"`
}

func SaveBlog(blog Blog) error {
	if blog.Title == "" || blog.Type == "" {
		return fmt.Errorf("invalid blog data")
	}

	db := config.GetDB()

	if blog.BlogID == "" {
		// create blog_id
		uuid := ""
		for i := 0; i < 10; i++ {
			tmp := utils.GetUUID()
			var cnt int64
			err := db.Table(tableBlog).Where("blog_id=?", tmp).Count(&cnt).Error
			if err != nil {
				return err
			}
			if cnt == 0 {
				uuid = tmp
				break
			}
		}
		if uuid == "" {
			return fmt.Errorf("create uuid error")
		}
		blog.BlogID = uuid

		// check parent id valid
		if blog.ParentID != "" {
			exist, err := ExistBlog(blog.ParentID)
			if err != nil {
				return err
			}
			if !exist {
				blog.ParentID = ""
			}
		}
		return db.Table(tableBlog).Create(blog).Error
	}

	return config.GetDB().Table(tableBlog).Save(blog).Error
}

func ExistBlog(blogID string) (bool, error) {
	var cnt int64
	err := config.GetDB().Table(tableBlog).Where("blog_id=?", blogID).Count(&cnt).Error
	return cnt == 1, err
}

func GetBlog(blogID string) (Blog, error) {
	blog := Blog{}
	err := config.GetDB().Table(tableBlog).Where("blog_id=?", blogID).First(&blog).Error
	return blog, err
}

type BlogBrief struct {
	BlogID   string `gorm:"column:blog_id"`
	Title    string `gorm:"column:title"`
	ParentID string `gorm:"column:parent_id"`
}

func GetAllBlogBrief() ([]BlogBrief, error) {
	blogList := make([]BlogBrief, 0)
	err := config.GetDB().Table(tableBlog).Find(&blogList).Error
	return blogList, err
}
