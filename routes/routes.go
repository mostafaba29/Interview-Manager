package routes

import (
	"mostafaba29/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Post("/auth/managerappoints", handlers.ShowManagerAppointments)
	app.Post("/auth/create", handlers.CreateAppointment)
	app.Patch("auth/manager/approve/:id", handlers.Approve)
	app.Patch("auth/manager/decline/:id", handlers.CancelAppointment)
	app.Patch("auth/manager/update/:id", handlers.UpdateAppointment)
	app.Get("/allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("/showbyclient", handlers.ShowClient)

}
