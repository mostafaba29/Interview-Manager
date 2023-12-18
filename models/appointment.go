package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	MeetingTime string `gorm:"not null" json:"meeting_time"`
	UserID      uint   `json:"user_id"`
	ManagerName string `json:"manager_name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Client      string `json:"client"`
	User        User
}
