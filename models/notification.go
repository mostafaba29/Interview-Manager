package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Text   string
	UserID uint
	User   User
}
