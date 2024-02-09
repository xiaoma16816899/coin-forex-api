package mdware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	fiberLimiter "github.com/gofiber/fiber/v2/middleware/limiter"
	"server.com/platform/cache"
)

func Limiter(max int, ttl time.Duration, custom ...func(c *fiberLimiter.Config) error) func(*fiber.Ctx) error {
	config := fiberLimiter.Config{
		Max:        max,
		Expiration: ttl,
		Storage:    cache.GetStore(),
	}

	if len(custom) > 0 {
		custom[0](&config)
	}

	return fiberLimiter.New(config)
}
