package router

import (
	"github.com/gofiber/fiber/v2"
	userRoute "server.com/pkg/router/authRoutes"
	crudRoute "server.com/pkg/router/crudRoutes"
	"server.com/pkg/router/forexRoutes"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	forexRoutes.SetupForexRouter(api)
	crudRoute.SetupTradingCURDRouter(api)
	userRoute.SetupUserRouter(api)

}
