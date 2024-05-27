package models

import "gorm.io/gorm"

// Perfume model
type Perfume struct {
	gorm.Model
	name  string
	Price int
}
