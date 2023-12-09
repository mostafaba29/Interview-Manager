package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateAppointment(c *fiber.Ctx) {

	var appointmentDetails models.Appointment
	if err := c.BodyParser(&appointmentDetails); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could't read appointment details",
		})

	}

	appointmentDetails = models.Appointment{
		ManagerID:   appointmentDetails.ManagerID,
		MeetingTime: appointmentDetails.MeetingTime,
		UserID:      appointmentDetails.UserID,
		Description: appointmentDetails.Description,
		Status:      "pending",
	}

	appointmentResult := intialization.DB.Create(&appointmentDetails).Error
	if appointmentResult.Error != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could't create appointemnt",
		})
	}

	c.JSON(fiber.Map{
		"massege": "appointment created",
	})

}

func CancelAppointment(c *fiber.Ctx) error {
	appointmentID := c.Params("id")
	var canceledAppointment models.Appointment
	if err := intialization.DB.First(&canceledAppointment, appointmentID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "appointment not found",
		})
	}
	canceledAppointment.Status = "canceled"
	intialization.DB.Save(&canceledAppointment)
	return c.JSON(fiber.Map{
		"massege": "appointment canceled",
	})
}

func UpdateAppointment(c *fiber.Ctx) error {

	var oldAppointment models.Appointment
	oldAppointmentID := c.Params("id")
	if err := intialization.DB.First(&oldAppointment, oldAppointmentID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"massege": "appointment not found",
		})
	}
	var updatedAppointment models.Appointment
	if err := c.BodyParser(&updatedAppointment); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could not update appointment",
		})
	}
	oldAppointment.Description = updatedAppointment.Description
	oldAppointment.MeetingTime = updatedAppointment.MeetingTime
	intialization.DB.Save(&oldAppointment)

	return c.Status(200).JSON(fiber.Map{
		"massege": "appointment updated",
	})

}
