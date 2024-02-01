package configs

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsConfig() cors.Config {
	// Define server settings.
	return cors.Config{
		// AllowOrigins: os.Getenv("ALLOW_ORIGINS"),
		// AllowHeaders: "Origin, Content-Type, Accept",
		AllowHeaders:     "*",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}
}
 