package models

import "gorm.io/gorm"

// Laptop model
type Laptop struct {
	gorm.Model
	Brand string
	Price int
}
