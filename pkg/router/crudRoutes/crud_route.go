package crudRoute

import (
	"github.com/gofiber/fiber/v2"
	tradingCRUDCtl "github.com/xiaoma/trading/app/Controller/tradingCRUD"
)

func SetupTradingCURDRouter(router fiber.Router) {
	route := router.Group("/trading")
	route.Post("/get_trading_list", tradingCRUDCtl.GetTradingList)
	route.Post("/add", tradingCRUDCtl.CreateTrading)
	route.Post("/update", tradingCRUDCtl.UpdateTrading)
	route.Post("/get_trading", tradingCRUDCtl.GetSingleTradingByID)
	route.Post("/delete", tradingCRUDCtl.DeleteForex)

}
