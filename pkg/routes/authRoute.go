package routes

import (
	"github.com/JacobRWebb/InventoryManagement/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

func AuthRoute(a *fiber.App) {
	authRoute := a.Group("/auth")
	loginRoute := authRoute.Group("/login")

	loginRoute.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	loginRoute.Post("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	})

	registerRoute := authRoute.Group("/register")

	registerRoute.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotImplemented)
	})

	registerRoute.Post("/", controller.RegisterAccount)
}
