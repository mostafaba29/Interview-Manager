package handlers

import (
	"mostafaba29/intialization"
	"mostafaba29/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func CreateAppointment(c *fiber.Ctx) error {

	var appointmentDetails models.Appointment
	if err := c.BodyParser(&appointmentDetails); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could not read appointment details",
		})

	}

	appointmentDetails = models.Appointment{
		ManagerName: appointmentDetails.ManagerName,
		MeetingTime: appointmentDetails.MeetingTime,
		UserID:      appointmentDetails.UserID,
		Description: appointmentDetails.Description,
		Client:      appointmentDetails.Client,
		Status:      "pending",
	}

	if err := intialization.DB.Create(&appointmentDetails).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could't create appointemnt",
		})
	}

	return c.JSON(fiber.Map{
		"massege":     "appointment created",
		"appointment": appointmentDetails,
	})

}

func Approve(c *fiber.Ctx) error {
	var appointment models.Appointment
	result := intialization.DB.First(&appointment, c.Params("id"))
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Appointment not found"})
	}

	intialization.DB.Model(&models.Appointment{}).Where("id = ?", c.Params("id")).Update("status", "Confirmed")

	return c.JSON(appointment)
}

func CancelAppointment(c *fiber.Ctx) error {
	appointmentID := c.Params("id")
	var canceledAppointment models.Appointment
	if err := intialization.DB.First(&canceledAppointment, appointmentID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "appointment not found",
		})
	}
	intialization.DB.Model(&models.Appointment{}).Where("id = ?", appointmentID).Update("status", "Declined")
	intialization.DB.Save(&canceledAppointment)
	return c.Status(200).JSON(canceledAppointment)
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
		"massege":  "appointment updated",
		"new_time": updatedAppointment,
	})

}
