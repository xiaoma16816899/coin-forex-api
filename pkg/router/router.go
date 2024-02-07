package router

import (
	"github.com/gofiber/fiber/v2"
	userRoute "server.com/pkg/router/authRoutes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	userRoute.SetupUserRouter(api)
	userRoute.SetupCaptcha(api)

}
