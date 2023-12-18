package handlers

import (
	"fmt"
	"mostafaba29/intialization"
	"mostafaba29/middleware"
	"mostafaba29/models"

	"github.com/gofiber/fiber/v2"
)

func GetCurrentUser(c *fiber.Ctx) (*models.User, error) {
	sess, err := middleware.Store.Get(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from session: %v", err)
	}

	userID := sess.Get(middleware.USER_ID)
	if userID == nil {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "not autherized cant get id",
		})
	}

	var user models.User
	intialization.DB.Where("id=?", userID).First(&user)
	if err != nil {
		return nil, c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "not autherized something is wrong",
		})
	}

	return &user, c.Status(fiber.StatusOK).JSON(user)
}

func ShowAppointments(c *fiber.Ctx) error {
	var appointments []models.Appointment
	intialization.DB.Find(&appointments)
	return c.JSON(appointments)
}

func ShowManagerAppointments(c *fiber.Ctx) error {
	user, err := GetCurrentUser(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	appointments := []models.Appointment{}

	managerName := user.Username
	if err := intialization.DB.Preload("User").Where("manager_name = ?", managerName).Find(&appointments).Error; err != nil {
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
