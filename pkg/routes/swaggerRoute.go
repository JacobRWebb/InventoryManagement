package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(a *fiber.App) {
	route := a.Group("/swagger")

	route.Use(swagger.HandlerDefault)
}
