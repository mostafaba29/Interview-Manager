package intialization

import "mostafaba29/models"

func MigrateDB() {

	DB.AutoMigrate(&models.Manager{}, &models.Admin{}, &models.Appointment{}, &models.AppointmentDetails{}, &models.Client{}, &models.Notification{})
}
