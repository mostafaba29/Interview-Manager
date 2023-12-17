package handlers

import (
	"log"
	"mostafaba29/intialization"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
)

func ShowAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	intialization.DB.Find(&appointments)
	return c.JSON(appointments)
}

func ShowManagerAppointments(c *fiber.Ctx) error {
	var managerAppointments []models.Appointment
	var manager models.User
	if err := intialization.DB.Model(&models.Appointment{}).Where("manager_name=?", "manger"); err != nil {
		log.Println(manager.Username)
		return c.Status(400).JSON(fiber.Map{
			"massege": "no appointments found",
		})
	}
	return c.Status(200).JSON(managerAppointments)
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
