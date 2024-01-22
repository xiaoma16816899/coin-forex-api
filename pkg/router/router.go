package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/xiaoma/trading/pkg/router/forexRoutes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	forexRoutes.SetupForexRouter(api)

}
