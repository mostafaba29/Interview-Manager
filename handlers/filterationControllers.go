package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func ShowAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	intialization.DB.Find(&appointments)
	return c.JSON(appointments)
}

func ShowManagerAppointments(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, _ := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return PrivateKey, nil
	})

	claims := token.Claims.(jwt.MapClaims)
	manager, _ := claims["issuer"].(string)

	var appointments []models.Appointment
	if err := intialization.DB.Where("manager_name = ?", manager).Find(&appointments).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(appointments)
}

func ShowApprovedAppointments(c *fiber.Ctx) error {
	var approvedAppointments []models.Appointment
	if err := intialization.DB.Where("status=?", "Confirmed").First(&approvedAppointments); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"massege": "no approved appointments found",
		})
	}
	return c.JSON(approvedAppointments)
}

func ShowCanceledAppointments(c *fiber.Ctx) error {
	var canceledAppointments []models.Appointment
	if err := intialization.DB.Where("status=?", "Declined").First(&canceledAppointments); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"massege": "no canceled appointments found",
		})

	}
	return c.JSON(canceledAppointments)
}
