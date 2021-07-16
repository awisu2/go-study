package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)



func OpenSqlite() *gorm.DB{
	return open(sqlite.Open("test.db"))
}

func OpenPostgresql() *gorm.DB{
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	return open(postgres.Open(dsn))
}

func open(dialector gorm.Dialector) *gorm.DB{
	config := CreateGormConfig()
	db, err := gorm.Open(dialector, config);
	if err != nil {
		panic(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
  db.AutoMigrate(&Product{}, &User{}, &ROnetooneBase{}, &ROnetooneItem{}, &RHasmanyBase{}, &RHasmanyItem{}, &RMtomBase{}, &RMtomItem{})
}

func CreateGormConfig() *gorm.Config {
	// logger: https://gorm.io/ja_JP/docs/logger.html
	//
	// SlowThresholdは挙動を確認したいため0にセット。これを超えないと、Loglevelがinfoでも出力されない
	//
	gormLogger := logger.New (
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config {
			SlowThreshold: 0,  // スローログ判定時間(通常はtime.Secondなど)
			LogLevel: logger.Info, // 
			IgnoreRecordNotFoundError: true,
			Colorful: true, // 色付
		},
	)

	return &gorm.Config{
		Logger: gormLogger,
	}
}