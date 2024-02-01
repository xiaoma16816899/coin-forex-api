package configs

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	googleChat "server.com/pkg"
	"server.com/pkg/types"
	logger "server.com/pkg/utils/logger"
)

// https://docs.gofiber.io/api/middleware/recover
func RecoverConfig() recover.Config {

	return recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, panicErr interface{}) {
			jsonBytes := types.JSON(debug.Stack())
			if _, ok := panicErr.(*fiber.Error); !ok {
				// Log only when, panic by unexpected, eg: access non exist index of an array
				logger.Debug("[panic]: ", panicErr, ", [url]:", c.OriginalURL(), ", [stack]:", jsonBytes)
				go googleChat.GetGoogleChat().Send(googleChat.Payload{
					Title: "💥 *Crash Panic Error*",
					Message: "path: " + c.OriginalURL() + "\n\n" +
						jsonBytes.String(),
				})
			}
		},
	}
}
