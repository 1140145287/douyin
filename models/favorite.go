package models

import "gorm.io/plugin/soft_delete"

// Favorite 用户点赞关系实体
type Favorite struct {
	Id      int64 `gorm:"column:favorite_id;autoIncrement"`
	VideoId int64
	UserId  int64
	IsDel   soft_delete.DeletedAt `gorm:"softDelete:flag"` //soft delete 0/1
}
