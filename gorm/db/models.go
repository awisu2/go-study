package db

import "gorm.io/gorm"

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type User struct {
	// add id, created_at, updated_at, deleted_at
  // gorm.Model
  ID        uint           `gorm:"primaryKey"`
	// 実際のカラムでは user_id になる
  UserId string `gorm:"uniqueIndex,not null"`
	Name string
  CompanyRefer string
  Company   Company `gorm:"references:Code;foreignKey:CompanyRefer;constraint:OnUpdate:SET NULL,OnDelete:CASCADE;"`
}

type Company struct {
  ID   int `gorm:"uniqueIndex,not null"`
  Code string `gorm:"uniqueIndex,not null"`
  Name string
}

type ROnetooneBase struct {
  ID        uint           `gorm:"primaryKey"`

  // one to one
  // ROneのカラムは作成されない
  ROnetooneItem      ROnetooneItem `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ROnetooneItem struct {
  ID        uint           `gorm:"primaryKey"`
  // RBaseのROneと連動(default: テーブル名+ID)
  ROnetooneBaseID   uint
  Name string `gorm:"required"`
}

type RHasmanyBase struct {
  ID  uint `gorm:"primaryKey"`
  Num int
  RHasmanyItems []RHasmanyItem
}

type RHasmanyItem struct {
  ID  uint `gorm:"primaryKey"`
  Num int
  RHasmanyBaseID uint
}

// many to many
// https://gorm.io/ja_JP/docs/many_to_many.html
//
// この場合 r_mtom_base_r_mtom_items という中間テーブルが作成される
// (中間テーブルはmany2manyの指定に準拠)
type RMtomBase struct {
  ID  uint `gorm:"primaryKey"`
  Num int
  RMtomItems []*RMtomItem `gorm:"many2many:RMtomBase_RMtomItems;"`
}

type RMtomItem struct {
  ID  uint `gorm:"primaryKey"`
  Num int
  RMtomBases []*RMtomBase  `gorm:"many2many:RMtomBase_RMtomItems;"`
}
