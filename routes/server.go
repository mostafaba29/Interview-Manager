package routes

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	Store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func Setup() {
	app := fiber.New()
	SetupRoutes(app)
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 24 * 30,
	})
	//app.Use(middleware.AuthMiddleware())
	log.Fatal(app.Listen(":8000"))
}
