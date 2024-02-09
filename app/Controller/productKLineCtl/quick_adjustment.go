package productKLineCtl

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"server.com/app/models"
	"server.com/app/service"
)

type targetParams struct {
	Fluctuation  float64 `json:"fluctuation" validate:"required"`
	CurrentPrice float64 `json:"current_price" validate:"required"`
	TargetPrice  float64 `json:"target_price" validate:"required"`
	Speed        int     `json:"speed" validate:"required"`
}

func GetTargetList(c *fiber.Ctx) error {
	params := new(targetParams)
	if err := c.BodyParser(params); err != nil {
		return models.FailWithMessage(models.INVALID_JSON, err.Error(), c)
	}
	klines := service.GenerateDataTradingView(params.Fluctuation, params.Speed, params.TargetPrice, params.CurrentPrice)
	data := make(map[string]interface{})
	data["klines"] = klines
	data["count"] = len(klines)
	return models.OkWithData(data, c)
}

func GenerateKListDataTrading(c *fiber.Ctx) error {
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 2, 1, 1, 30, 0, 0, time.UTC) // Example: End date after 30 minutes
	currentPrice := 2568.00
	fluctuationPercentage := 0.05 // Fluctuation percentage (e.g., 5%)
	upProbability := 0.7          // Probability of an up movement
	downProbability := 0.3        // Probability of a down movement
	tradingData := service.GenerateTradingDataS(startDate, endDate, currentPrice, fluctuationPercentage, upProbability, downProbability)
	return models.OkWithData(tradingData, c)
}
