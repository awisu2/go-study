package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TryGorm() {
	// db := openSqlite()
	db := openPostgresql()

  // Migrate the schema
  db.AutoMigrate(&Product{}, &User{})

  // Create
  db.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.First(&product, 1) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)

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