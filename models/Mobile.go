package models

import "gorm.io/gorm"

// Mobile model
type Mobile struct {
	gorm.Model
	Brand string
	Price int
	ElectronicID int
	Electronic Electronic `gorm:"foreignKey:ElectronicID;references:ID"`
}
