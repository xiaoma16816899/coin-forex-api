package authCtl

import (
	"github.com/gofiber/fiber/v2"
	"server.com/app/models"
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
	return models.OkWithDetailed(users, "操作成功", c)
}

func GetUser(c *fiber.Ctx) error {
	json := new(userParams)
	if err := c.BodyParser(json); err != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	user, err := service.GetUser(json.UserID)

	if err != nil {
		return models.FailWithMessage(-100000, err.Error(), c)
	}

	return models.OkWithDetailed(user, "Get successful", c)
}

func CreateUser(c *fiber.Ctx) error {

	json := new(models.User)
	if ok := c.BodyParser(json); ok != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	hash, _ := service.HashPassword(json.Password)
	json.Password = hash

	ok := service.CreateUser(*json)
	if ok != nil {
		return models.FailWithMessage(-50000, "internal server error", c)
	}

	return models.OkWithMessage("create new user is successfully", c)
}

func UpdateUser(c *fiber.Ctx) error {
	json := new(models.User)
	if err := c.BodyParser(json); err != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	ok := service.UpdateUser(*json)
	if ok != nil {
		return models.FailWithMessage(-50000, "Internal Server Error", c)
	}

	return models.OkWithMessage("update user successfully", c)
}

func DeleteUser(c *fiber.Ctx) error {
	json := new(userParams)
	if ok := c.BodyParser(json); ok != nil {
		return models.FailWithMessage(-400, "Invalid Json", c)
	}

	user, ok := service.GetUser(json.UserID)

	if ok != nil {
		return models.FailWithMessage(-10000, ok.Error(), c)
	}

	ok = service.DeleteUser(int64(user.ID))
	if ok != nil {
		return models.FailWithMessage(-50000, ok.Error(), c)
	}
	return models.FailWithMessage(0, "Delete successfully", c)

}
