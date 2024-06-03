package models

import "gorm.io/gorm"

// Pen model
type Pen struct {
	gorm.Model
	Name string
}
