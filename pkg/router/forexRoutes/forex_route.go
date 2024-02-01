package forexRoutes

import (
	"github.com/gofiber/fiber/v2"
	controller "server.com/app/Controller/forexCtl"
)

func SetupForexRouter(router fiber.Router) {
	route := router.Group("/forex")
	route.Post("/get_exchange_currency", controller.GetForexExChange)
	route.Post("get_all_exchange_currency", controller.GetAllForexExChange)

}
