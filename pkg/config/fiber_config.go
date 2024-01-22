package configs

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

// FiberConfig func for configuration Fiber app.
// See: https://docs.gofiber.io/api/fiber#config
func FiberConfig() fiber.Config {
	// Define server settings.
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	return fiber.Config{
		BodyLimit:   3 * 1024 * 1024, //3mb
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		ErrorHandler: FiberErrorHandler,
	}
}

func FiberErrorHandler(c *fiber.Ctx, err error) error {
	resErr := fiber.Error{
		Code:    fiber.ErrInternalServerError.Code,
		Message: fiber.ErrInternalServerError.Message,
	}

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		resErr.Code = e.Code
		resErr.Message = e.Message
	}

	if err != nil {
		return c.Status(resErr.Code).JSON(resErr)
	}

	// Return from handler
	return nil
}
