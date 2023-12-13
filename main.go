package main

import (
	"log"
	"mostafaba29/intialization"
	"mostafaba29/middleware"
	"mostafaba29/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	intialization.ConnectToDB()
	intialization.MigrateDB()
}

func main() {
	app := fiber.New()

	app.Use(middleware.AuthMiddleware())
	app.Use(middleware.CorsMiddleware())

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8000"))
}
