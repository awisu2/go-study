package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 一度開いたDBを保管
var instanceDBs map[string]Database = map[string]Database{}

type Database struct {
  DB *gorm.DB
}

func Open() Database {
  key := "default"
  return OpenByKey(key)
}

func OpenByKey(key string) Database {
  database, ok := instanceDBs[key]

  // db open
  if !ok {
    db := openDB("sqlite")
    database = Database{DB: db}
    instanceDBs[key] = database
  }

  return database
}

// 実際にDBを開く処理
func openDB(kind string) *gorm.DB {
  log.Println("open DB. " + kind)

  if kind == "postgresql" {
    return openPostgresql()
  } else {
    return openSqlite()
  }
}

func openSqlite() *gorm.DB{
  if db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}); err == nil {
		return db
	} else {
		panic(err)
	}
}

func openPostgresql() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
  if db, err := gorm.Open(postgres.Open(dsn)); err == nil {
		return db
	} else {
		panic(err)
	}
}

