package models

import "gorm.io/gorm"

// Food model
type Food struct {
	gorm.Model
	Name string
	Type string
	IsCheap string
}
