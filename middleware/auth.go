package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	Store    *session.Store
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func AuthMiddleware() fiber.Handler {
	return NewMiddleware
}

func NewMiddleware(c *fiber.Ctx) error {
	Store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSameSite: fiber.CookieSameSiteLaxMode,
		Expiration:     time.Hour * 24,
	})

	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}

	session, err := Store.Get(c)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"massege": "unautherized",
		})
	}

	if session.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"massege": "unautherized to proceed",
		})
	}
	return c.Next()
}
