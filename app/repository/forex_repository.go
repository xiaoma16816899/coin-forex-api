package repositories

import (
	"gorm.io/gorm"
)

// UserQueries struct for queries from User model.
type Forex struct {
	db *gorm.DB
}

func AdminUserRepo(db *gorm.DB) *Forex {
	return &Forex{db}
}
