package models

import "gorm.io/gorm"

// Monitor model
type Monitor struct {
	gorm.Model
	Brand string
}
