package repositories

import (
	"gorm.io/gorm"
)

// UserQueries struct for queries from User model.
type Trading struct {
	db *gorm.DB
}

func TradingRepo(db *gorm.DB) *Trading {
	return &Trading{db}
}
