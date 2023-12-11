package routes

import (
	"mostafaba29/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Post("/create", handlers.CreateAppointment)
	app.Patch("approve", handlers.Approve)
	app.Patch("/decline", handlers.CancelAppointment)
	app.Patch("/update", handlers.UpdateAppointment)
	app.Get("allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("showbyclient", handlers.ShowClient)

}
