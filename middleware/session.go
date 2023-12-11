package middleware

import (
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

	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSameSite: fiber.CookieSameSiteLaxMode,
		Expiration:     time.Hour * 24,
	})
}
