package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewMiddleware() fiber.Handler {
	return AuthMiddleware
}
func AuthMiddleware(c *fiber.Ctx) error {
	session, err := Store.Get(c)
	if strings.Split(c.Path(), "/")[1] == "auth" {
		return c.Next()
	}
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
