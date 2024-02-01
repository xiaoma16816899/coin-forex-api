package service

import (
	"server.com/app/model"
	"server.com/platform/database"
)

func GetUserPaginated(page int, limit int) ([]model.User, error) {
	db, _ := database.OpenDBConnection()
	var users []model.User
	offset := page * limit
	sql := "SELECT * FROM user WHERE "

	sql = sql + " u.deleted_at IS NULL ORDER BY u.created_at DESC LIMIT ? OFFSET ? ;"
	res := db.Raw(sql, limit, offset).Scan(&users)

	if res.RowsAffected == 0 {
		return make([]model.User, 0), nil
	}
	return users, nil
}

func GetUser(userId int) (model.User, error) {
	db, _ := database.OpenDBConnection()
	var user model.User
	sql := "SELECT * FROM user WHERE "
	sql = sql + " u.deleted_at IS NULL AND id = ? ;"
	res := db.Raw(sql, userId).Find(&user)

	if res.RowsAffected != 0 {
		return user, res.Error
	}

	return user, nil
}

func CreateUser(user model.User) error {
	db, _ := database.OpenDBConnection()
	ok := db.Save(&user).Error
	return ok
}

func UpdateUser(user model.User) error {
	db, _ := database.OpenDBConnection()
	ok := db.Save(&user).Error
	return ok
}

func DeleteUser(userID int64) error {
	db, _ := database.OpenDBConnection()
	err := db.Where("id = ?", userID).Delete(&model.User{}).Error
	return err
}
