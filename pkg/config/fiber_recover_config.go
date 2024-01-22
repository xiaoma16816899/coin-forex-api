package configs

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	googleChat "github.com/xiaoma/trading/pkg"
	"github.com/xiaoma/trading/pkg/types"
	logger "github.com/xiaoma/trading/pkg/utils"
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
