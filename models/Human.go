package models

import "gorm.io/gorm"

// Human model
type Human struct {
	gorm.Model
	IsFemale int
	AnimalID int
	Animal Animal `gorm:"foreignKey:AnimalID;references:ID"`
}
