package models

import "gorm.io/gorm"

// Game model
type Game struct {
	gorm.Model
	Name string
	YearFounded string
}
