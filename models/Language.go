package models

import "gorm.io/gorm"

// Language model
type Language struct {
	gorm.Model
	Name      string
	IsAncient int
}
