package main

import (
	"log"
	"os"

	"github.com/JacobRWebb/InventoryManagement/pkg/database"
	"github.com/JacobRWebb/InventoryManagement/pkg/middleware"
	"github.com/JacobRWebb/InventoryManagement/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	database.InitDB()

	defer database.DB.Close()

	middleware.FiberMiddleware(app)

	// routes.SwaggerRoute(app)
	routes.AuthRoute(app)
	routes.UserRoute(app)
	routes.RouteNotFound(app)

	app.Listen(":" + os.Getenv("PORT"))

}
