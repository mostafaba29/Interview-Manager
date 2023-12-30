package middleware

import (
	"fmt"
	"mostafaba29/handlers"
	"mostafaba29/intialization"
	"mostafaba29/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return RequireAuth
}

func RequireAuth(c *fiber.Ctx) error {
	tokenString := c.Cookies("jwt")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return handlers.PrivateKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token signature",
			})
		}
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.Status(fiber.StatusUnauthorized)
		}

		var user models.User
		intialization.DB.First(&user, claims["issuer"])

		if user.ID == 0 {
			c.Status(fiber.StatusUnauthorized)
		}
		c.Locals("user", user)
	}
	return c.Next()
}
