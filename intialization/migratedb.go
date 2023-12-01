package intialization

import "mostafaba29/models"

func Migrate() {

	DB.AutoMigrate(&models.Admin{}, &models.Appointment{}, &models.AppointmentDetails{}, &models.Client{}, &models.Manager{}, &models.Notification{})
}
