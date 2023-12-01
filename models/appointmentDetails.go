package models

import (
	"time"

	"gorm.io/gorm"
)

type AppointmentDetails struct {
	gorm.Model
	Approved      bool
	Deleted       bool
	Updated       bool
	NewTime       time.Time
	AppointmentID uint
	Appointment   Appointment
}
