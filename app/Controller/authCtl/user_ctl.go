package authCtl

import (
	"github.com/gofiber/fiber/v2"
	"server.com/app/model"
	"server.com/app/service"
	"server.com/pkg/pagination"
)

type userParams struct {
	UserID int `json:"userId"`
}

func GetUsers(c *fiber.Ctx) error {
	// payload := struct {
	// 	Username string `json:"username"`
	// }{}
	// c.BodyParser(&payload)
	page, limit := pagination.GetPaginationParams(c)
	users, _ := service.GetUserPaginated(page, limit)
	return model.OkWithDetailed(users, "操作成功", c)
}

func GetUser(c *fiber.Ctx) error {
	json := new(userParams)
	if err := c.BodyParser(json); err != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}

	user, err := service.GetUser(json.UserID)

	if err != nil {
		return model.FailWithMessage(-100000, err.Error(), c)
	}

	return model.OkWithDetailed(user, "Get successful", c)
}

func CreateUser(c *fiber.Ctx) error {

	json := new(model.User)
	if ok := c.BodyParser(json); ok != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}

	hash, _ := service.HashPassword(json.PasswordHash)
	json.PasswordHash = hash

	ok := service.CreateUser(*json)
	if ok != nil {
		return model.FailWithMessage(-50000, "internal server error", c)
	}

	return model.OkWithMessage("create new user is successfully", c)
}

func UpdateUser(c *fiber.Ctx) error {
	json := new(model.User)
	if err := c.BodyParser(json); err != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}

	ok := service.UpdateUser(*json)
	if ok != nil {
		return model.FailWithMessage(-50000, "Internal Server Error", c)
	}

	return model.OkWithMessage("update user successfully", c)
}

func DeleteUser(c *fiber.Ctx) error {
	json := new(userParams)
	if ok := c.BodyParser(json); ok != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}

	user, ok := service.GetUser(json.UserID)

	if ok != nil {
		return model.FailWithMessage(-10000, ok.Error(), c)
	}

	ok = service.DeleteUser(int64(user.ID))
	if ok != nil {
		return model.FailWithMessage(-50000, ok.Error(), c)
	}
	return model.FailWithMessage(0, "Delete successfully", c)

}
