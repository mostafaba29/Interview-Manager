package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Text                 string
	AppointmentDetailsID uint
	AppointmentDetails   AppointmentDetails
	ManagerID            uint
	Manager              Manager
}
