package dbs

import (
	"time"

	"gorm.io/gorm"
)

// テーブルの宣言は一箇所にまとめて管理
func AutoMigrate() {
  db := Open("")
  db.DB.AutoMigrate(&Product{}, &User{})
}

// gorm.Modelの拡張
// jsonレスポンスの設定ができないため、明示する
type GormModel struct {
  ID        uint `gorm:"primarykey" json:"id"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
  DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted"`
}

// 実際に作成されるテーブルは products とs がつく
type Product struct {
  // gorm.Model により ID, CreatedAt, UpdatedAt, DeletedAt を含む
  gorm.Model
  Code  string
  Price uint
}

type User struct {
  GormModel
  UserId string `gorm:"unique" json:"userId"`
  Name string   `json:"name"`
  AccessedAt time.Time
}

