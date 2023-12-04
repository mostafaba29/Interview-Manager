package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string
	Password      []byte
	Position      string
	Appointments  []Appointment
	Notifications []Notification
}
