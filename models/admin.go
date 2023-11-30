package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	name     string
	password string
	Manager  Manager
}
