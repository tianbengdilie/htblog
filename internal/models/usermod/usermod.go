// package usermod
package usermod

import (
	"htblog/internal/config"
	mhelper "htblog/internal/models/helper"
)

const (
	tableUser = "t_users"
)

type UserInfo struct {
	mhelper.Model
	Account  string `gorm:"column:account"`
	Password string `gorm:"column:password"`
	NickName string `gorm:"column:nickname"`
}

// CreateUser
func CreateUser(account, pwd, NickName string) error {
	db := config.GetDB()
	return db.Table(tableUser).Create(&UserInfo{
		Account:  account,
		Password: pwd,
		NickName: NickName,
	}).Error
}

// GetUser
func GetUser(account, pwd string) (*UserInfo, error) {
	db := config.GetDB()
	var user UserInfo
	err := db.Table(tableUser).Where("account=? and password=?", account, pwd).First(&user).Error
	if mhelper.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}
