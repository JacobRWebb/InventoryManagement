package application

import (
	"os"

	"github.com/JacobRWebb/InventoryManagement/pkg/store/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type App struct {
	App *fiber.App
	DB  *gorm.DB
}

func MustCreateNewApp() *App {
	app := fiber.New(fiber.Config{
		// Views: engine,
	})

	db := database.MustOpen(os.Getenv("DATABASE_NAME"))

	return &App{
		App: app,
		DB:  db,
	}
}
