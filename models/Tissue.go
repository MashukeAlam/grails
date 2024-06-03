package models

import "gorm.io/gorm"

// Tissue model
type Tissue struct {
	gorm.Model
	Name string
}
