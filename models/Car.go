package models

import "gorm.io/gorm"

// Car model
type Car struct {
	gorm.Model
	Name string
	Type string
	IsCheap string
}
