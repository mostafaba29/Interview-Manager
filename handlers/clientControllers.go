package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
)

func ShowClient(c *fiber.Ctx) error {

	name := c.Query("client")
	var clients models.Appointment
	intialization.DB.Where("client LIKE ?", "%"+name+"%").First(&clients)
	return c.JSON(clients)
}
