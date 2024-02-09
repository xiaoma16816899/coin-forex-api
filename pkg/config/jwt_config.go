package configs

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JwtConfig func for configuration jwt auth.
// See: https://github.com/gofiber/jwt
func JwtConfig() jwtware.Config {
	return jwtware.Config{
		SigningKey:    []byte(os.Getenv("JWT_SECRET_KEY")),
		SigningMethod: jwt.SigningMethodHS256.Name,
		ContextKey:    "jwt", // used in private routes
		ErrorHandler:  jwtError,
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing or malformed JWT",
		})
	}

	return c.Status(416).JSON(fiber.Map{
		"code": 416,
		"msg":  "无效 Token",
	})
}
