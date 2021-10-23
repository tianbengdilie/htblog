package blogmod

import (
	"htblog/internal/config"
	mhelper "htblog/internal/models/helper"

	"gorm.io/gorm"
)

type Blog struct {
	mhelper.Model
	BlogID  string `gorm:"column:blog_id"`
	Title   string `gorm:"column:title"`
	Type    string `gorm:"column:type"`
	Content string `gorm:"column:content"`
	Editor  string `gorm:"column:editor"`
}

const tableBlog = "t_blogs"

func SaveBlog(blog Blog) error {
	db := config.GetDB()

	if blog.BlogID == "" {
		// create blog_id
		return db.Table(tableBlog).Create(blog).Error
	}

	dbBLog := &Blog{}
	db.Table(tableBlog).Where("blog")

	return config.GetDB().Table(tableBlog).Save(blog).Error
}

func GetBlog(blogID string) (*Blog, error) {
	blog := &Blog{}
	err := config.GetDB().Table(tableBlog).Where("blog_id=?", blogID).First(blog).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return blog, nil
}
