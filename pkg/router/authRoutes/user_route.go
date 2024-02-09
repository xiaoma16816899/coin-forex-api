package userRoute

import (
	"github.com/gofiber/fiber/v2"
	"server.com/app/Controller/user"
)

func SetupUserRouter(router fiber.Router) {
	route := router.Group("/user")
	route.Post("/get_list", user.GetUsers)
	route.Post("/add", user.CreateUser)
	route.Post("/update", user.DeleteUser)
	route.Post("/get", user.GetUser)
	route.Post("/delete", user.DeleteUser)
	route.Post("/get_send_code", user.SendCodeMobileVerify)
}
