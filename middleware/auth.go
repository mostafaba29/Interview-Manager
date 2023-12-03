package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return Authenticate
}

func Authenticate(c *fiber.Ctx) error {
	session, err := Store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "login" {
		return c.Next()
	}

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not Autherized",
		})
	}

	if session.Get(AUTH_KEY) == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Not Autherized",
		})
	}

	return c.Next()
}
