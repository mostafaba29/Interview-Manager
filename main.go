package main

import (
	"log"
	"mostafaba29/intialization"

	"github.com/gofiber/fiber/v2"
)

func init() {
	intialization.ConnectToDB()
	intialization.GormMigrate()
	intialization.LoadDotEnv()
}

func home(c *fiber.Ctx) error {
	return c.SendString("de7k")

}
func main() {
	app := fiber.New()

	app.Get("/home", home)

	log.Fatal(app.Listen("PORT"))
}
