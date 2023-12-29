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
<<<<<<< HEAD
	app.Patch("/manager/approve/:id", handlers.Approve)
	app.Patch("/manager/decline/:id", handlers.CancelAppointment)
=======
	app.Patch("/auth/manager/approve/:id", handlers.Approve)
	app.Patch("/auth/manager/decline/:id", handlers.CancelAppointment)
>>>>>>> 5d0d2e4f8d2653787a27480df57ff777e0689d4b
	app.Patch("/manager/update/:id", handlers.UpdateAppointment)
	app.Get("/auth/allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("/showbyclient", handlers.ShowClient)
<<<<<<< HEAD
	//app.Get("/auth/getuser", handlers.GetCurrentUser)
=======
>>>>>>> 5d0d2e4f8d2653787a27480df57ff777e0689d4b
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
