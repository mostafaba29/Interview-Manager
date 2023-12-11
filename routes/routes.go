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
	app.Post("approve", handlers.Approve)
	app.Post("/decline", handlers.CancelAppointment)
	app.Post("/update", handlers.UpdateAppointment)
	app.Get("allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("showbyclient", handlers.ShowClient)

}
