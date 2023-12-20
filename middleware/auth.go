package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	Store *session.Store
	User  string = "username"
)

func AuthMiddleware() fiber.Handler {
	return NewMiddleware
}

func NewMiddleware(c *fiber.Ctx) error {

	session, err := Store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"massege": "unautherized error in session",
		})
	}

	if session.Get(User) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"massege": "unautherized to proceed",
		})
	}
	return c.Next()
}
