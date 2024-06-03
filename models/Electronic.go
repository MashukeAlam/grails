package models

import "gorm.io/gorm"

// Electronic model
type Electronic struct {
	gorm.Model
	Type string
	Name string
	Inventor string
}
