package authTypes

import (
	"fmt"
)

type Response struct {
	Token string `json:"token"`
	User  any    `json:"user"`
}

type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}
type SignUp struct {
	Username string `json:"email" validate:"required,username,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	Role     string `json:"role"`
	Name     string `json:"name" validate:"lte=25"`
}

// SignIn struct to describe login user.
type LogIn struct {
	Username string `json:"email" validate:"required,username,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
}

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Username string `json:"username"`
	UserID   uint   `json:"userId"`
	Exp      int64  `json:"exp"`
	// Permisions  map[types.Permision]bool `json:"permisions"`
	IpWhiteList string `json:"ip_white_list"`
	GrantType   string `json:"grantType"`
	// role id
	RoleID uint `json:"RoleID"`
}

func (t TokenMetadata) Valid() error {
	if t.UserID == 0 {
		return fmt.Errorf("invalid email")
	}
	return nil
}
