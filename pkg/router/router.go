package router

import (
	"github.com/gofiber/fiber/v2"
	userRoute "server.com/pkg/router/authRoutes"
	"server.com/pkg/router/productKLineRoutes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	userRoute.SetupUserRouter(api)
	userRoute.SetupCaptcha(api)
	productKLineRoutes.SetupKLineProductRouter(api)
}
