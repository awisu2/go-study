package dbs

func TryDB() {
	db := Open("")

  // Migrate the schema
  db.DB.AutoMigrate(&Product{}, &User{})

  // Create
  db.DB.Create(&Product{Code: "D42", Price: 100})

  // Read
  var product Product
  db.DB.First(&product, 1) // find product with integer primary key
  db.DB.First(&product, "code = ?", "D42") // find product with code D42

  // Update - update product's price to 200
  db.DB.Model(&product).Update("Price", 200)
  // Update - update multiple fields
  db.DB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.DB.Delete(&product, 1)
}
