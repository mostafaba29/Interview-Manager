package routes

import (
	"mostafaba29/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Post("/auth/managerappoints", handlers.ShowManagerAppointments)
	app.Post("/create", handlers.CreateAppointment)
	app.Patch("/manager/approve/:id", handlers.Approve)
	app.Patch("/manager/decline/:id", handlers.CancelAppointment)
	app.Patch("/manager/update/:id", handlers.UpdateAppointment)
	app.Get("allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("showbyclient", handlers.ShowClient)

}
