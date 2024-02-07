package repositories

import (
	"gorm.io/gorm"
)

// UserQueries struct for queries from User model.
type Admin struct {
	db *gorm.DB
}

func AdminRepo(db *gorm.DB) *Admin {
	return &Admin{db}
}
