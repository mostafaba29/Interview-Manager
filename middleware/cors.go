package middleware

import "github.com/gofiber/fiber/v2"

func CorsMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Allow credentials (cookies, HTTP authentication)
		c.Set("Access-Control-Allow-Credentials", "true")

		// Specify allowed origins (replace with your actual frontend domain)
		c.Set("Access-Control-Allow-Origin", "http://localhost:5500")

		// Specify allowed methods
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

		// Specify allowed headers
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Continue to the next middleware/handler in the chain
		return c.Next()
	}
}
