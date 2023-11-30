package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	name    string
	purpose string
}
