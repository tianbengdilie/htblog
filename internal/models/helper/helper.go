package helper

import (
	"errors"
	"time"

	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint      `gorm:"column:id;primarykey;autoIncrement"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func IsNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}

func IsUniqueConstraintError(err error) bool {
	return errors.Is(err, sqlite3.ErrConstraintUnique)
}

func IsPrimaryConstraintError(err error) bool {
	return errors.Is(err, sqlite3.ErrConstraintPrimaryKey)
}
