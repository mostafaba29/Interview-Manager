package routes

import (
	"mostafaba29/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/auth/logout", handlers.Logout)
	app.Post("/auth/manager/:username", handlers.ShowManagerAppointments)
	app.Post("/auth/create", handlers.CreateAppointment)
	app.Patch("/manager/approve/:id", handlers.Approve)
	app.Patch("/manager/decline/:id", handlers.CancelAppointment)
	app.Patch("/manager/update/:id", handlers.UpdateAppointment)
	app.Get("/auth/allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("/showbyclient", handlers.ShowClient)
	app.Get("/auth/getbyname/:username", handlers.GetUserByUsername)
	app.Get("/auth/getuser", handlers.GetCurrentUser)

}
