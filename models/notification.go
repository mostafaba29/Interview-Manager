package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	text string
}
