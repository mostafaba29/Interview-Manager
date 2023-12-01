package routes

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {

	app.Get("/home", home)

}

func home(c *fiber.Ctx) error {
	return c.SendString("de7k")

}
