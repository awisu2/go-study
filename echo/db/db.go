package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func Open() *gorm.DB{
	if db == nil {
		var err error
		db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	return db
}

// Migrate the schema
func Migrate(db *gorm.DB) {
	// migrationファイルなどの作成はされず、追加や変更にのみ対応(migration処理は別ライブラリでやったほうが良い)
	db.AutoMigrate(&Product{})
}