package main

import (
	"log"
	"mostafaba29/intialization"
	"mostafaba29/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func init() {
	intialization.ConnectToDB()
	intialization.MigrateDB()
	intialization.LoadDotEnv()
}

func main() {
	app := fiber.New()
	store := session.New(session.ConfigDefault)
	routes.SetupRoutes(app)
	app.Use(store)
	log.Fatal(app.Listen(":8000"))

}
