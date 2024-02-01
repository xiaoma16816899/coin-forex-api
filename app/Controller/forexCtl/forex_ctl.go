package controller

import (
	"github.com/gofiber/fiber/v2"
	model "server.com/app/model"
	"server.com/openApi"
)

func GetForexExChange(c *fiber.Ctx) error {
	params := new(model.ParamForexTradingModel)
	if err := c.BodyParser(params); err != nil {
		return model.FailWithMessage(-10000, "Invalid Json", c)
	}

	result, err := openApi.GetForexExchangeCurrencies(*params)
	if err != nil {
		return model.ErrorResult(-1, err.Error(), c)
	}

	return model.OkWithData(result, c)
}

func GetAllForexExChange(c *fiber.Ctx) error {
	type filterParam struct {
		Date string `json:"date"`
	}
	params := new(filterParam)
	if err := c.BodyParser(params); err != nil {
		return model.FailWithMessage(-10000, "Invalid Json", c)
	}

	result, err := openApi.GetAllForexExchangeCurrencies(params.Date)
	if err != nil {
		return model.ErrorResult(-1, err.Error(), c)
	}

	return model.OkWithData(result, c)
}
