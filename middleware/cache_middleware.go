package mdware

import (
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	fiberCache "github.com/gofiber/fiber/v2/middleware/cache"
	"server.com/platform/cache"
)

func reqCacheKeyGenerator(c *fiber.Ctx) string {
	return cache.PrefixReq + c.OriginalURL()
}

// Use Minimal configs
func UseCache(ttl time.Duration) fiber.Handler {
	handler := fiberCache.New(fiberCache.Config{
		Storage: cache.GetStore(),
		Next: func(c *fiber.Ctx) bool {
			if strings.Contains(c.Get(fiber.HeaderCacheControl), "no-cache") {
				return true
			}
			return c.Query("refresh") == "true"
		},
		Expiration:   ttl,
		KeyGenerator: reqCacheKeyGenerator,
	})

	return handler
}

func SkipCache() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Request().Header.Set(fiber.HeaderCacheControl, "no-cache")
		return c.Next()
	}
}

// Client side caching
func CSC(ttl time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		maxAge := strconv.FormatUint(uint64(ttl.Seconds()), 10)
		c.Set(fiber.HeaderCacheControl, "public, max-age="+maxAge)
		return c.Next()
	}
}

// Use full options
// https://docs.gofiber.io/api/middleware/cache
func Cache(config fiberCache.Config) fiber.Handler {

	// Use redis storage
	config.Storage = cache.GetStore()
	config.KeyGenerator = reqCacheKeyGenerator
	handler := fiberCache.New(config)

	return handler
}
