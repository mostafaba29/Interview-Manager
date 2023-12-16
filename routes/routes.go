package routes

import (
	"mostafaba29/handlers"
	"mostafaba29/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)
	app.Post("/logout", middleware.AuthMiddleware(), handlers.Logout)
	app.Post("/managerappoints", middleware.AuthMiddleware(), handlers.ShowManagerAppointments)
	app.Post("/auth/create", middleware.AuthMiddleware(), handlers.CreateAppointment)
	app.Patch("/manager/approve/:id", middleware.AuthMiddleware(), handlers.Approve)
	app.Patch("/manager/decline/:id", middleware.AuthMiddleware(), handlers.CancelAppointment)
	app.Patch("/manager/update/:id", middleware.AuthMiddleware(), handlers.UpdateAppointment)
	app.Get("/auth/allappointments", handlers.ShowAppointments)
	app.Get("/showapporved", handlers.ShowApprovedAppointments)
	app.Get("/showdeclined", handlers.ShowCanceledAppointments)
	app.Get("/showbyclient", handlers.ShowClient)

}
