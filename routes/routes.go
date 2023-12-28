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
	// app.Post("/auth/create", handlers.CreateAppointment)
	// app.Patch("/manager/approve/:id", handlers.Approve)
	// app.Patch("/manager/decline/:id", handlers.CancelAppointment)
	// app.Patch("/manager/update/:id", handlers.UpdateAppointment)
	app.Get("/auth/allappointments", handlers.ShowAppointments)
	// app.Get("/showapporved", handlers.ShowApprovedAppointments)
	// app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	// app.Get("/showbyclient", handlers.ShowClient)
	// app.Get("/auth/getuser", handlers.GetCurrentUser)
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
