package routes

import (
	"github.com/JacobRWebb/InventoryManagement/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(a *fiber.App) {
	route := a.Group("/user")

	route.Post("/GetUserByEmail", controller.GetUserByEmail)
}
