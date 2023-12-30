package handlers

import (
	"fmt"
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
	appointment := models.Appointment{
		MeetingTime: appointmentDetails.MeetingTime,
		Client:      appointmentDetails.Client,
		Description: appointmentDetails.Description,
		ManagerName: appointmentDetails.ManagerName,
		Status:      "Pending",
	}

	intialization.DB.Model(&models.User{}).Association("Appointment").Append(appointment)
	if err := intialization.DB.Create(&appointment).Error; err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could't create appointemnt",
		})
	}

	return c.Status(fiber.StatusOK).JSON(appointment)

}

func Approve(c *fiber.Ctx) error {
	var appointment models.Appointment
	result := intialization.DB.First(&appointment, c.Params("id"))
	fmt.Println(c.Params("id"))
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
	return c.Status(200).JSON(canceledAppointment)
}

func UpdateAppointment(c *fiber.Ctx) error {

	var UpdatedAppointment models.Appointment
	UpdatedAppointmentID := c.Params("id")
	if err := intialization.DB.First(&UpdatedAppointment, UpdatedAppointmentID).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"massege": "appointment not found",
		})
	}
	var NewAppointment models.Appointment
	if err := c.BodyParser(&NewAppointment); err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "could not update appointment",
		})
	}
	UpdatedAppointment.Description = NewAppointment.Description
	UpdatedAppointment.MeetingTime = NewAppointment.MeetingTime
	intialization.DB.Save(&UpdatedAppointment)

	return c.Status(200).JSON(UpdatedAppointment)
}
