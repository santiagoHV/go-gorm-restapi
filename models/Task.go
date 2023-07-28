package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model

	Id          uint   `gorm:"primaryKey"` // Set field as primary key
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Done        bool   `gorm:"default:false"`
	UserID      uint   `gorm:"not null"`
}
