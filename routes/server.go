package routes

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/mattn/go-sqlite3"
)

var (
	Store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func Setup() {
	app := fiber.New()
	SetupRoutes(app)

	storage := sqlite3.New()
	store := session.New(session.Config{
		Storage: storage,
	})

	app.Use(session.New(session.Config{
		KeyLookup:      "header:Authorization",
		Storage:        store,
		CookieName:     "auth",
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 24 * 30,
	}))
	log.Fatal(app.Listen(":8000"))
}
