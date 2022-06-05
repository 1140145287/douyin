package models

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

//Comment 存放用户评论的实体
type Comment struct {
	Id         int64                 `json:"id,omitempty"`
	User       User                  `json:"user" sql:"-" gorm:"-"` //不入库，只做前端返回
	UserId     int64                 `json:"user_id"`
	VideoId    int64                 `json:"video_id"`
	Content    string                `json:"content,omitempty"`
	CreatedAt  time.Time             `json:"created_at,omitempty"`
	CreateDate string                `json:"create_date" sql:"-" gorm:"-"` //ignore 不入库
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag"`              //soft delete 0/1
}
