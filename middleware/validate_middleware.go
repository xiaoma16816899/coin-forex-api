package mdware

import (
	"github.com/gofiber/fiber/v2"
	"server.com/pkg/types"
	"server.com/pkg/utils"
)

// Validate only body of the request
func ValidateBody(structure interface{}) fiber.Handler {
	validator := func(c *fiber.Ctx) error {
		if err := c.BodyParser(structure); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})

		}

		isValid, errors := utils.ValidatorErrors(structure)
		if !isValid {
			return c.Status(fiber.StatusBadRequest).JSON(types.ValidateErrorResponse{
				Code:    10011,
				Message: "Validation error",
				Errors:  errors,
			})
		}

		return c.Next()
	}

	return validator
}
