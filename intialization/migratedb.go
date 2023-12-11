package intialization

import "mostafaba29/models"

func MigrateDB() {

	DB.AutoMigrate(&models.User{}, &models.Appointment{}, &models.Notification{})
}
