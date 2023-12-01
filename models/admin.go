package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	ManagerID uint
	Manager   Manager
}
