package models

type Session struct {
	ID     string `gorm:"primaryKey"`
	UserID uint
	User   User
	Data   string
}
