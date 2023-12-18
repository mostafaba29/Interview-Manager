package handlers

import (
	"fmt"
	"mostafaba29/intialization"
	"mostafaba29/middleware"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) error {
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return fmt.Errorf("failed to get user from session: %v", err)
	}

	userID := sess.Get(middleware.USER_ID)
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "not autherized cant get id",
		})
	}

	var user models.User
	intialization.DB.Where("id=?", userID).First(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "not autherized something is wrong",
		})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func ShowAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	intialization.DB.Find(&appointments)
	return c.JSON(appointments)
}

func ShowManagerAppointments(c *fiber.Ctx) error {
	username := c.Params("username")
	var managerAppointments []models.Appointment

	if err := intialization.DB.Where("manager_name = ?", username).Find(&managerAppointments).Error; err != nil {
		return c.Status(404).SendString("no appointments found" + err.Error())
	}
	return c.Status(200).JSON(managerAppointments)
}

func GetUserByUsername(c *fiber.Ctx) error {
	username := c.Params("username")

	var user models.User
	if err := intialization.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}

	return c.JSON(user)
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
