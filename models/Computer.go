package models

import "gorm.io/gorm"

// Computer model
type Computer struct {
	gorm.Model
	Brand string
}
