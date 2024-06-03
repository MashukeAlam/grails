package models

import "gorm.io/gorm"

// Girl model
type Girl struct {
	gorm.Model
	SexualPreference string
}
