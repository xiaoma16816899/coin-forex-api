package captChaCtl

import (
	"bytes"
	"image/png"

	"github.com/gofiber/fiber/v2"
	"server.com/app/models"
	"server.com/app/service"
)

func ImageVerify(c *fiber.Ctx) error {
	randomString := service.GenerateRandomString(service.NumChars)
	image, err := service.CreateImage(randomString)
	if err != nil {
		return models.FailWithMessage(models.INTERNAL_SERVER_ERROR, err.Error(), c)
	}

	var imgBuffer bytes.Buffer
	if err := png.Encode(&imgBuffer, image); err != nil {
		return models.FailWithMessage(models.INTERNAL_SERVER_ERROR, err.Error(), c)

	}

	c.Context().SetContentType("image/png")
	if _, err := c.Write(imgBuffer.Bytes()); err != nil {
		return models.FailWithMessage(models.INTERNAL_SERVER_ERROR, err.Error(), c)

	}
	return nil

}
