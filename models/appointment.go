package models

import (
	"time"

	"gorm.io/gorm"
)

type Appointment struct {
	gorm.Model
	meetTime time.Time
	Admin    Admin
	Manager  Manager
	reason   string
}
