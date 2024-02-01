package userRoute

import (
	"github.com/gofiber/fiber/v2"
	"server.com/app/Controller/authCtl"
)

func SetupUserRouter(router fiber.Router) {
	route := router.Group("/user")
	route.Post("/get_list", authCtl.GetUsers)
	route.Post("/add", authCtl.CreateUser)
	route.Post("/update", authCtl.DeleteUser)
	route.Post("/get", authCtl.GetUser)
	route.Post("/delete", authCtl.DeleteUser)
}
