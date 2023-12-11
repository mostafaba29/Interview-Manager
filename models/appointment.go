package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	MeetingTime time.Time `gorm:"not null" json:"meeting_time"`
	UserID      uint      `json:"user_id"`
	ManagerID   uint      `json:"manager_id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Client      string    `json:"client"`
}
