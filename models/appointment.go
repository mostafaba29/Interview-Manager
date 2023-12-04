package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	MeetingTime time.Time `gorm:"not null"`
	UserID      uint
	ManagerID   uint
	Description string
}
