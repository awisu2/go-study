package db

import (
	"go-study/gorm/libs"
)



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
