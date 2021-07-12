package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



func OpenSqlite() *gorm.DB{
  if db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}); err == nil {
		return db
	} else {
		panic(err)
	}
}

func OpenPostgresql() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  if db, err := gorm.Open(postgres.Open(dsn)); err == nil {
		return db
	} else {
		panic(err)
	}
}

func AutoMigrate(db *gorm.DB) {
  db.AutoMigrate(&Product{}, &User{})
}