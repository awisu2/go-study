package main

import (
	"fmt"
	"go-study/gorm/db"

	"gorm.io/gorm"
)


func main() {
	// db := database.OpenSqlite()
	DB := db.OpenPostgresql()

  db.AutoMigrate(DB)

  sampleGetstart(DB)
  sampleCreateUser(DB)
  sampleRelation(DB)
}

func sampleGetstart(DB *gorm.DB) {
  // Create
  DB.Create(&db.Product{Code: "D42", Price: 100})

  // Read
  var product db.Product
  DB.First(&product, 1) // find product with integer primary key
  DB.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  DB.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  DB.Model(&product).Updates(db.Product{Price: 200, Code: "F42"}) // non-zero fields
  DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  DB.Delete(&product, 1)
}

func sampleCreateUser(DB *gorm.DB) {
	DB.Create(db.CreateUser("taro"))
}

func sampleRelation(DB *gorm.DB) {
  fmt.Println("sampleRelation")
}
