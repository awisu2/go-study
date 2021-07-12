package database

import (
	"go-study/gorm/libs"

	"gorm.io/gorm"
)

type User struct {
	// add id, created_at, updated_at, deleted_at
  gorm.Model
	// 実際のカラムでは user_id になる
  UserId string `gorm:"uniqueIndex,not null"`
	Name string
}

func CreateUser(name string) *User {
	userId, err := libs.CreateRandomString(20, libs.RANDOM_LETTERS)
	if err != nil {
		panic(err)
	}

	return &User{
		UserId: userId,
		Name: name,
	}
}
