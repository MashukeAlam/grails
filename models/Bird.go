package models

import "gorm.io/gorm"

// Bird model
type Bird struct {
	gorm.Model
	Name string
	IsAncient int
	AnimalID int
	Animal Animal `gorm:"foreignKey:AnimalID;references:ID"`
}
