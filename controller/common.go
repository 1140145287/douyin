package controller

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title"`
}

type Comment struct {
	Id         int64                 `json:"id,omitempty"`
	User       User                  `json:"user" sql:"-" gorm:"-"`
	UserId     int64                 `json:"user_id"`
	VideoId    int64                 `json:"video_id"`
	Content    string                `json:"content,omitempty"`
	CreatedAt  time.Time             `json:"created_at,omitempty"`
	CreateDate string                `json:"create_date" sql:"-" gorm:"-"` //ignore
	IsDel      soft_delete.DeletedAt `gorm:"softDelete:flag"`              //soft delete 0/1
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:user_id" `
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id;autoIncrement"`
	VideoId int64
	UserId  int64
	IsDel   soft_delete.DeletedAt `gorm:"softDelete:flag"` //soft delete 0/1
}
