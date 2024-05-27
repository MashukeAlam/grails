package models

import "gorm.io/gorm"

// Mouse model
type Mouse struct {
	gorm.Model
	Brand string
	Price int
}
