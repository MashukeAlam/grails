// Package internals Never TOUCH this file please.
package internals

import (
	"gorm.io/gorm"
	"github.com/MashukeAlam/grails/models"
)

func Migrate(db *gorm.DB) {
	// add models to watch for migration.
	db.AutoMigrate(&models.Game{})
	db.AutoMigrate(&models.Player{})
	db.AutoMigrate(&models.Country{})
	db.AutoMigrate(&models.Food{})
	db.AutoMigrate(&models.Car{})
	db.AutoMigrate(&models.Bike{})
	db.AutoMigrate(&models.Animal{})
	db.AutoMigrate(&models.Human{})
	db.AutoMigrate(&models.Language{})
	db.AutoMigrate(&models.Bird{})
	db.AutoMigrate(&models.Electronic{})
	db.AutoMigrate(&models.Girl{})
	db.AutoMigrate(&models.Laptop{})
	db.AutoMigrate(&models.Monitor{})
	db.AutoMigrate(&models.Pen{})
	db.AutoMigrate(&models.Tissue{})
	db.AutoMigrate(&models.Perfume{})
	db.AutoMigrate(&models.Mobile{})
	db.AutoMigrate(&models.Smartphone{})
	db.AutoMigrate(&models.Card{})
	db.AutoMigrate(&models.Mouse{})
}
