package models

import "gorm.io/gorm"

type Product struct {
  // gorm.Model により ID, CreatedAt, UpdatedAt, DeletedAt を含む
  gorm.Model
  Code  string
  Price uint
}

type User struct {
  gorm.Model
  UserId string `gorm:"unique"`
  Name string
}

