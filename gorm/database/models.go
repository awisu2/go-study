package database

import "gorm.io/gorm"

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type User struct {
	// add id, created_at, updated_at, deleted_at
  gorm.Model
	// 実際のカラムでは user_id になる
  UserId  string `gorm:"uniqueIndex,not null"`
	Name string
}
