package models

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
  DB *gorm.DB
}

type DbKind string
const (
  DB_KIND_SQLITE = DbKind("sqlite")
  DB_KIND_POSTGRESQL = DbKind("postgresql")
)
type Config struct {
  Key string
  Kind DbKind
}

type Configs map[string]Config

var _configs Configs
var instanceDBs map[string]Database = map[string]Database{}

func SetConfigs(configs Configs)  {
  _configs = configs
}

func Open() Database {
  key := "default"
  return OpenByKey(key)
}

func OpenByKey(key string) Database {
  database, ok := instanceDBs[key]

  // db open
  if !ok {
    config := _configs[key]
    db := openDB(config.Kind)
    database = Database{DB: db}
    instanceDBs[key] = database
  }

  return database
}

// 実際にDBを開く処理
func openDB(kind DbKind) *gorm.DB {
  log.Println("open DB. " + kind)

  switch kind {
  case "postgresql":
    return openPostgresql()
  default:
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

