package repositories

import (
	"gorm.io/gorm"
)

type Admin struct {
	db *gorm.DB
}

func AdminRepo(db *gorm.DB) *Admin {
	return &Admin{db}
}
