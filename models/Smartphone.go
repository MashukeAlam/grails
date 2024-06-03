package models

import "gorm.io/gorm"

// Smartphone model
type Smartphone struct {
	gorm.Model
	Brand string
	Price int
}
