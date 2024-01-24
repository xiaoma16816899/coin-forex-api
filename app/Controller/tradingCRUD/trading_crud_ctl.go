package tradingCRUDCtl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xiaoma/trading/app/model"
	"github.com/xiaoma/trading/app/service"
)

type params struct {
	ID uint64 `json:"id"`
}

func GetTradingList(c *fiber.Ctx) error {
	result, err := service.GetTradings()
	if err != nil {
		return model.FailWithMessage(-5000, err.Error(), c)
	}
	return model.OkWithData(result, c)
}

func GetSingleTradingByID(c *fiber.Ctx) error {
	json := new(params)

	if err := c.BodyParser(json); err != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}

	result, err := service.GetTradingByID(json.ID)
	if err != nil {
		return model.FailWithMessage(-5000, err.Error(), c)
	}
	return model.OkWithData(result, c)
}

func CreateTrading(c *fiber.Ctx) error {
	json := new(model.ForexTradingModel)

	if err := c.BodyParser(json); err != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}
	err := service.UpdateForex(json)

	if err != nil {
		return model.FailWithMessage(-50000, err.Error(), c)
	}
	return model.OkWithMessage("Delete successfully", c)
}

func UpdateTrading(c *fiber.Ctx) error {
	json := new(model.ForexTradingModel)
	if err := c.BodyParser(json); err != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}
	err := service.UpdateForex(json)
	if err != nil {
		return model.FailWithMessage(-50000, err.Error(), c)
	}

	return model.OkWithMessage("Update successfully", c)
}

func DeleteForex(c *fiber.Ctx) error {
	json := new(params)
	if err := c.BodyParser(json); err != nil {
		return model.FailWithMessage(-400, "Invalid Json", c)
	}
	err := service.DeleteForex(json.ID)
	if err != nil {
		return model.FailWithMessage(-50000, err.Error(), c)
	}

	return model.OkWithMessage("Delete successfully", c)
}
