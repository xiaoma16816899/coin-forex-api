package model

import "time"

type User struct {
	BaseModel
	PasswordHash   string     `json:"-"`
	Name           string     `json:"name"`
	Nickname       string     `json:"nickname"`
	Username       string     `json:"username" gorm:"unique" validate:"required"`
	VerifyCode     string     `json:"codeVerify"`
	CurrentLoginIP string     `json:"currentLoginIP"`
	LastLoginIP    string     `json:"lastLoginIP"`
	CurrentRegion  string     `json:"currentRegion"`
	LastRegion     string     `json:"lastRegion"`
	IsActive       bool       `json:"isActive" gorm:"default:true"`
	CurrentLoginAt *time.Time `json:"currentLoginAt" sql:"DEFAULT:NULL"`
	LastLoginAt    *time.Time `json:"lastLoginAt" sql:"DEFAULT:NULL"`
	Balance        float64    `json:"balance"  gorm:"default:0"`
}
