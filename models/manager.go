package models

import "gorm.io/gorm"

type Manager struct {
	gorm.Model
	name     string
	password string
	//Appointment Appointment
}
