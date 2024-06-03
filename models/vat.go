package models

import "gorm.io/gorm"

// Vat model
type Vat struct {
	gorm.Model
	Name string
}
