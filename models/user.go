package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null; unique" json:"username"`
	Password []byte `gorm:"not null" json:"password"`
	Position string `gorm:"not null" json:"position"`
}
