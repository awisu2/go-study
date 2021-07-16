package db

import "gorm.io/gorm"

type Product struct {
  gorm.Model
  Code  string
  Price uint
}

type User struct {
	// add id, created_at, updated_at, deleted_at
  gorm.Model
	// 実際のカラムでは user_id になる
  UserId string `gorm:"uniqueIndex,not null"`
	Name string
}


type RBase struct {
  ID        uint           `gorm:"primaryKey"`
  // belongs to (キーとセットで記述 defaultだと名前＋ID)
  //
  // DBにカラムは生成されない
  //
  RParent    RParent
  RParentID  uint
  //
  // 外部キーを変えたい場合
  // RParent    RParent `gorm:"foreignKey:RParentRefer"`
  // RParentRefer  uint
  //
  // 関連するキー対象を変えたい場合(相手のキーはuniqueであること)
  // RParent    RParent `gorm:"references:SubID"`
}

type RParent struct {
  ID        uint           `gorm:"primaryKey"`
  SubID     uint           `gorm:"unique"`
}



type ROne struct {
  ID        uint           `gorm:"primaryKey"`
  
}