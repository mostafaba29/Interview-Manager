package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
)

func ShowAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	intialization.DB.Find(&appointments)
	return c.JSON(appointments)
}

func ShowApprovedAppointments(c *fiber.Ctx) error {
	var approvedAppointments []models.Appointment
	if err := intialization.DB.Where("status=?", "approved").First(&approvedAppointments); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"massege": "no approved appointments found",
		})
	}
	return c.JSON(approvedAppointments)
}

func ShowCanceledAppointments(c *fiber.Ctx) error {
	var canceledAppointments []models.Appointment
	if err := intialization.DB.Where("status=?", "canceled").First(&canceledAppointments); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"massege": "no canceled appointments found",
		})

	}
	return c.JSON(canceledAppointments)
}
