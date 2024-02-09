package productKLineRoutes

import (
	"github.com/gofiber/fiber/v2"
	"server.com/app/Controller/productKLineCtl"
)

func SetupKLineProductRouter(router fiber.Router) {
	route := router.Group("/productKline")
	route.Post("/get_target_list", productKLineCtl.GetTargetList)
	route.Post("/get_k_list", productKLineCtl.GenerateKListDataTrading)
}
