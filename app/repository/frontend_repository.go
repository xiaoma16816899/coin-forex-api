package repositories

import (
	"gorm.io/gorm"
)

type Frontend struct {
	db *gorm.DB
}

func FrontendRepo(db *gorm.DB) *Frontend {
	return &Frontend{db}
}
