package main

import (
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}

	app

	// db := database.MustOpen(os.Getenv("DATABASE_NAME"))
	// passwordHash := passwordhash.NewPasswordHash()

	// userStore := dbstore.NewUserStore(dbstore.NewUserStoreParams{
	// 	DB:           db,
	// 	PasswordHash: passwordHash,
	// })

	// engine := html.New("./web/views", ".templ")

	// app := fiber.New(fiber.Config{
	// 	Views: engine,
	// })

	// dependant := &Dependant{
	// 	app: app,
	// 	db:  db,
	// 	stores: *&Stores{
	// 		userStore: *userStore,
	// 	},
	// }

	// middleware.FiberMiddleware(app)

	// routes.SwaggerRoute(app)
	// routes.AuthRoute(app)
	// routes.UserRoute(app)
	// routes.RouteNotFound(app)

	// app.Listen(":" + os.Getenv("PORT"))
}
