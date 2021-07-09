package main

import (
	"go-study/gorm/database"
)


func main() {
	// db := database.OpenSqlite()
	db := database.OpenPostgresql()

  database.AutoMigrate(db)

  // Create
  db.Create(&database.Product{Code: "D42", Price: 100})

  // Read
  var product database.Product
  db.First(&product, 1) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.Model(&product).Updates(database.Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 1)

	db.Create(database.CreateUser("taro"))
}


