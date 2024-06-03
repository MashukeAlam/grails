package models

import "gorm.io/gorm"

// Bike model
type Bike struct {
	gorm.Model
	Name string
	Type string
	IsCheap int
}
