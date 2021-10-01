package db

import "gorm.io/gorm"


type Product struct {
	gorm.Model
	Code  string  `gorm:"uniqueIndex"`
	Price uint
}
