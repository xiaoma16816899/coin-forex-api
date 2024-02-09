package configs

import (
	"os"

	"github.com/gofiber/storage/redis"
	"server.com/pkg/utils"
)

func FiberRedisConfig() redis.Config {
	redisConnURL, _ := utils.ConnectionURLBuilder("redis")
	reset := false

	if os.Getenv("APP_ENV") == "stage" {
		reset = true
	}
	return redis.Config{
		URL:   redisConnURL,
		Reset: reset,
	}
}
