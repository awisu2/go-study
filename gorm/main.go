package main

import (
	"fmt"
	"go-study/gorm/db"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


func main() {
	// db := database.OpenSqlite()
	DB := db.OpenPostgresql()

  db.AutoMigrate(DB)

  sampleGetstart(DB)
  sampleCreateUser(DB)
  sampleRelationOneToOne(DB)
  sampleRelationHasMany(DB)
  sampleRelationMtom(DB)
}

func sampleGetstart(DB *gorm.DB) {
  fmt.Println("---------- sampleGetstart")
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
  fmt.Println("---------- sampleCreateUser")
	DB.Create(db.CreateUser("taro"))
}

func sampleRelationOneToOne(DB *gorm.DB) {
  fmt.Println("---------- sampleRelation")

  // one to oneの作成
  rOnetooneItem := db.ROnetooneItem{Name: "item"}
  rOnetooneBase := db.ROnetooneBase{ROnetooneItem: rOnetooneItem}
  DB.Create(&rOnetooneBase)
  log.Println(rOnetooneBase, rOnetooneBase.ROnetooneItem)

  // 単にselectするだけでは一緒に検索されない
  fmt.Println("-- only select")
  var rOnetooneBase2 db.ROnetooneBase
  if err := DB.First(&rOnetooneBase2, rOnetooneBase.ID).Error; err != nil {
    log.Println(err)
  }
  log.Println(rOnetooneBase2)

  // preload をすることで連携する値もセットされる
  // https://gorm.io/ja_JP/docs/preload.html
  //
  // 必要な分を事前に in 句 で検索している
  //
  fmt.Println("-- with preload")
  var rOnetooneBase3 db.ROnetooneBase
  if err := DB.Preload("ROnetooneItem").First(&rOnetooneBase3, rOnetooneBase.ID).Error; err != nil {
    log.Println(err)
  }
  log.Println(rOnetooneBase3)

  // 必要な関連テーブルすべてをpreload
  fmt.Println("-- with preload assosiations")
  if err := DB.Preload(clause.Associations).First(&rOnetooneBase3, rOnetooneBase.ID).Error; err != nil {
    log.Println(err)
  }
  log.Println(rOnetooneBase3)

  // joins をすることで連携する値もセットされる
  // 必要な分を検索時に Left Join 句 で検索している
  //
  // 基本的にPreloadの方で良いと思われる
  //
  fmt.Println("-- with joins")
  var rOnetooneBase4 db.ROnetooneBase
  if err := DB.Joins("ROnetooneItem").First(&rOnetooneBase4, rOnetooneBase.ID).Error; err != nil {
    log.Println(err)
  }
  log.Println(rOnetooneBase4)
  // DB.Create(&db.ROnetooneBase{})

}

func sampleRelationHasMany(DB *gorm.DB) {
  fmt.Println("---------- sampleRelationHasMany")
  items := []db.RHasmanyItem{
    {Num: 1},
    {Num: 2},
  }
  base := db.RHasmanyBase{Num: 3}
  base.RHasmanyItems = items
  if err := DB.Create(&base).Error; err != nil {
    log.Print(err)
  }

  fmt.Println("-- find preload")
  var bases []db.RHasmanyBase
  if err := DB.Preload(clause.Associations).Find(&bases).Error; err != nil {
    log.Println(err)
  }
  log.Println(bases)
}


func sampleRelationMtom(DB *gorm.DB) {
  fmt.Println("---------- sampleRelationMtom")
  items := []*db.RMtomItem{
    {Num: 1},
    {Num: 2},
  }
  base := db.RMtomBase{Num: 3}
  base.RMtomItems = items
  if err := DB.Create(&base).Error; err != nil {
    log.Print(err)
  }

  fmt.Println("-- append")
  DB.Model(&base).Association("RMtomItems").Append([]db.RMtomItem{
    {Num: 99},
  })
  fmt.Println("-- select after append")
  var base2 db.RMtomBase
  if err := DB.Preload(clause.Associations).First(&base2, base.ID).Error; err != nil {
    fmt.Println(err)
  } else {
    // 3つになっている
    for _, v := range base2.RMtomItems {
      fmt.Printf("item ID: %d, Num: %d\n", v.ID, v.Num)
    }
  }

  fmt.Println("-- find preload")
  var bases []db.RMtomBase
  if err := DB.Preload(clause.Associations).Find(&bases).Error; err != nil {
    log.Println(err)
  }
  log.Println(bases)

  // 中身の入ったレコードを元に何らかの処理を行う
  // https://gorm.io/ja_JP/docs/associations.html
  fmt.Println("-- association select")
  var items2 []db.RMtomItem
  if err :=DB.Model(&base).Association("RMtomItems").Find(&items2); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(items2)
  }
}
