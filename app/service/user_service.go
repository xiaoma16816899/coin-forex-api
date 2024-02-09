package service

import (
	"server.com/app/models"
	"server.com/platform/database"
)

func SubmitVerifyCode(verifyCode models.VerifyCode) error {
	db, _ := database.OpenDBConnection()
	ok := db.Save(&verifyCode).Error
	return ok
}
