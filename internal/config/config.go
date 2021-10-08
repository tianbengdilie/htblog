package config

import (
	"htblog/internal/pkg/tcache"
	"htblog/internal/pkg/tconfig"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	_config *tconfig.TConfig

	_db     *gorm.DB
	_dbOnce sync.Once

	_cacheUser *tcache.TCache
	_cacheOnce sync.Once
)

func Init() {
	_config = tconfig.New()
}

func GetDB() *gorm.DB {
	_dbOnce.Do(func() {
		dsn := "root:root@tcp(127.0.0.1:3306)/htblog?charset=utf8mb4&parseTime=True&loc=Local"
		var err error
		_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return _db
}

func GetUserCache() *tcache.TCache {
	_cacheOnce.Do(func() {
		_cacheUser = tcache.New()
	})
	return _cacheUser
}

func Config() *tconfig.TConfig {
	return _config
}
