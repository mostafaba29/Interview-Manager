package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `gorm:"not null"`
	Password      []byte `gorm:"not null"`
	Position      string `gorm:"not null"`
	Appointments  []Appointment
	Notifications []Notification
}
