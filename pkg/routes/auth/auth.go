package auth_routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthRoute struct {
	db *gorm.DB
}

func (a *AuthRoute) authRoute(app *fiber.App) {
	route := app.Group("/auth")

	route.Get("/register")
}
