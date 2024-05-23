package models

import "gorm.io/gorm"

// Country model
type Country struct {
	gorm.Model
	Name string
	Capital string
	Currency string
}
