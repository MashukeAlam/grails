package models

import "gorm.io/gorm"

// Player model
type Player struct {
	gorm.Model
	Name string
	GameID int
	Game Game `gorm:"foreignKey:GameID;references:ID"`
}
