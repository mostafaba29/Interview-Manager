package routes

import (
	"mostafaba29/handlers"
	"mostafaba29/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/signup", handlers.Signup)
	app.Post("/login", handlers.Login)
	app.Post("/auth/logout", middleware.AuthMiddleware(), handlers.Logout)
	app.Get("/auth/manager/appoints", middleware.AuthMiddleware(), handlers.ShowManagerAppointments)
	app.Post("/auth/create", middleware.AuthMiddleware(), handlers.CreateAppointment)
	app.Patch("/auth/manager/approve/:id", middleware.AuthMiddleware(), handlers.Approve)
	app.Patch("/auth/manager/decline/:id", middleware.AuthMiddleware(), handlers.CancelAppointment)
	app.Patch("auth/manager/update/:id", middleware.AuthMiddleware(), handlers.UpdateAppointment)
	app.Get("/auth/allappointments", middleware.AuthMiddleware(), handlers.ShowAppointments)
	app.Get("auth/showapporved", middleware.AuthMiddleware(), handlers.ShowApprovedAppointments)
	app.Get("auth/showdeclined", middleware.AuthMiddleware(), handlers.ShowCanceledAppointments)
	app.Get("auth/showbyclient", middleware.AuthMiddleware(), handlers.ShowClient)
	app.Options("/*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})
}
