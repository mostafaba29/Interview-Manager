package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	Name         string `gorm:"not null"`
	Password     string `gorm:"not null"`
	Appointment  []Appointment
	Notification []Notification
}
