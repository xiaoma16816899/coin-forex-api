package mdware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	configs "server.com/pkg/config"
)

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware
func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(configs.CorsConfig()),
		// Add simple logger.
		logger.New(logger.Config{
			Done: func(c *fiber.Ctx, logString []byte) {
				// AnalysisData(c)

				// make log
				// ActionLog(c)
			},
		}),
		// recover, when panic
		recover.New(configs.RecoverConfig()),
	)

}
