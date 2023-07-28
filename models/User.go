package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Id        uint   `gorm:"primaryKey"` // Set field as primary key
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Email     string `gorm:"not null" gorm:"unique"`
	Tasks     []Task `gorm:"foreignKey:UserID"`
}
