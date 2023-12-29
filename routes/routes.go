package routes

import (
	"mostafaba29/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Get("/auth/manager/appoints", handlers.ShowManagerAppointments)
	app.Post("/auth/create", handlers.CreateAppointment)
	app.Patch("/auth/manager/approve/:id", handlers.Approve)
	app.Patch("/auth/manager/decline/:id", handlers.CancelAppointment)
	app.Patch("auth/manager/update/:id", handlers.UpdateAppointment)
	app.Get("/auth/allappointments", handlers.ShowAppointments)
	app.Get("auth/showapporved", handlers.ShowApprovedAppointments)
	app.Get("auth/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("auth/showbyclient", handlers.ShowClient)
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
