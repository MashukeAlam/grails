package models

import "gorm.io/gorm"

// Animal model
type Animal struct {
	gorm.Model
	IsHuman  int
	PlayerID int
	Player   Player `gorm:"foreignKey:PlayerID;references:ID"`
}
