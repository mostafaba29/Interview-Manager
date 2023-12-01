package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	MeetingTime time.Time `gorm:"not null"`
	AdminID     uint
	ManagerID   uint
	Admin       Admin
	Manager     Manager
	Purpose     string
}
