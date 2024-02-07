package userRoute

import (
	"github.com/gofiber/fiber/v2"
	captChaCtl "server.com/app/Controller/captchaCtl"
)

func SetupCaptcha(router fiber.Router) {
	route := router.Group("/imageVerify")
	route.Get("/", captChaCtl.ImageVerify)
}
