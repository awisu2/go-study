package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
  gorm.Model
  Code  string  `gorm:"uniqueIndex"`
  Price uint
}

func main() {
  db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  migrate(db)
  crud(db)
}

// Migrate the schema
func migrate(db *gorm.DB) {
  // migrationファイルなどの作成はされず、追加や変更にのみ対応(migration処理は別ライブラリでやったほうが良い)
  db.AutoMigrate(&Product{})
}

func crud(db *gorm.DB) {
	// Create
	product := Product{Code: "D42", Price: 100}
	res := db.Create(&product)
	if err := res.Error; err != nil {
		log.Panic(err)
	}

	// Read
	var _product Product
	db.First(&_product, product.ID) // find product with integer primary key
	fmt.Println(_product)
	db.First(&_product, "code = ?", product.Code) // find product with code D42
	fmt.Println(_product)

	// Update - update product's price to 200
	db.Model(&_product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&_product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&_product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println(_product)

	// Delete - delete product
	db.Delete(&_product, product.ID)
	fmt.Println(_product)
}