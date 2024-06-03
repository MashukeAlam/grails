package models

import "gorm.io/gorm"

// Card model
type Card struct {
	gorm.Model
	Type string
	Face int
}
