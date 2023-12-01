package main

import (
	"log"
	"mostafaba29/intialization"
	"mostafaba29/routes"

	"github.com/gofiber/fiber/v2"
)

func init() {
	intialization.ConnectToDB()
	intialization.Migrate()
	intialization.LoadDotEnv()
}

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)
	log.Fatal(app.Listen("PORT"))
}
