package main

import (
	"fmt"
	"go-study/gorm/db"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)


func main() {
	DB := db.OpenSqlite()
	// DB := db.OpenPostgresql()

  db.AutoMigrate(DB)

  // sampleGetstart(DB)
  // sampleCreateUser(DB)
  // sampleRelationOneToOne(DB)
  // sampleRelationHasMany(DB)
  // sampleRelationMtom(DB)
  // SampleNullTime(DB)
  sampleRelation(DB)
  // sampleJoin(DB)
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


// nulltimeのwhereの挙動をテスト
//
// 通常のdatetimeフィールドと同じように扱っても良さそう
func SampleNullTime(DB *gorm.DB) {
  log.Println("---------- SampleNullTime")

  // NullTime に対して IS NOT NULLは有効
  users := []db.User{}
  DB.Where("updated_at IS NOT NULL").Find(&users)
  log.Println(len(users))
  DB.Where("updated_at IS NULL").Find(&users)
  log.Println(len(users))

  // NullTime に対して 時間検索
  DB.Where("updated_at < ?", time.Now()).Find(&users)
  log.Println(len(users))
  DB.Where("updated_at > ?", time.Now()).Find(&users)
  log.Println(len(users))
}

func sampleRelation(DB *gorm.DB) {

  printCompanies(DB)

  // 削除
  //
  // Unscoped(): 物理削除
  // .Session(&gorm.Session{AllowGlobalUpdate: true}): whereなしでもdelete
  //
  DB.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(db.Company{})
  DB.Unscoped().Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(db.User{})

  cp1 := db.Company{
    Name: "companyA",
    Code: "a",
  }
  cp2 := db.Company{
    Name: "companyB",
    Code: "b",
  }

  // 事前に宣言した値も更新したい場合は、参照にする必要あり(配列作成時にコピーされている)
  companies := []*db.Company {
    &cp1, &cp2,
  }
  if err:=DB.Create(companies).Error; err != nil {
    log.Panic(err)
  }

  // 注意: Companyは参照ではないため、Update,Save,Deleteの影響を受けない、扱いには注意
  _users := []db.User{
    {
      Name: "user1",
      Company: cp1,
    },
    {
      Name: "user2",
      Company: cp2,
    },
    {
      Name: "no company",
    },
  }
  if err:=DB.Create(_users).Error; err != nil {
    log.Panic(err)
  }

  log.Printf("ID: %d, name: %s, code: %s\n", cp1.ID, cp1.Name, cp1.Code)
  log.Printf("ID: %d, name: %s, code: %s\n", cp2.ID, cp2.Name, cp2.Code)
  for i := range companies {
    log.Printf("%v\n", companies[i])
  }

  for i := range _users {
    log.Printf("%v\n", _users[i])
  }


  var users []db.User
  DB.Find(&users)
  for i := range users {
    log.Println(users[i])
  }

  // ここでPreleadに指定しているのは、User structのフィールド名(テーブル名じゃない)
  DB.Preload("Company").Find(&users)
  for i := range users {
    log.Println(users[i])
  }

  sampleJoin(DB)

  // 更新
  cp1.Name = "cp1 update"
  if err:=DB.Save(&cp1).Error; err != nil {
    log.Panic(err)
  }
  if err:=DB.Model(&cp2).Updates(db.Company{Code: "bbb", Name: "cp2 update"}).Error; err != nil {
    log.Panic(err)
  }
  printCompanies(DB)
  printUsers(DB)

  // 削除
  if err:=DB.Unscoped().Delete(&cp2).Error; err != nil {
    log.Panic(err)
  }
  printCompanies(DB)
  printUsers(DB)

}

func printUsers(DB *gorm.DB) {
  var users []db.User
  DB.Preload("Company").Find(&users)
  log.Println("printUsers -----")
  for i := range users {
    user := users[i]

    log.Printf("%v: %s. Company: (%s) %v\n", user.ID, user.Name, user.CompanyRefer, user.Company)
  }
}

func printCompanies(DB *gorm.DB) {
  var users []db.Company
  DB.Find(&users)
  log.Println("printCompanies -----")
  for i := range users {
    log.Println(users[i])
  }
}

func sampleJoin(DB *gorm.DB) {
  log.Println("sampleJoin -----")
  var users []db.User
  res := DB.Model(&db.User{}).Select(
    "users.*",
    ).Joins(
      "INNER JOIN companies ON users.company_refer == companies.code",
    ).Where("companies.id > ?", 0).Scan(&users)
  if res.Error != nil {
    log.Panic(res.Error)
  }
  log.Println(&users)
}