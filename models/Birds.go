package models

import "gorm.io/gorm"

// Birds model
type Birds struct {
	gorm.Model
	Name string
	IsAncient int
}
