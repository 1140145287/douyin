package models

type Relation struct {
	Id          int64 `gorm:"column:relation_id"`
	FollowerId  int64 `gorm:"column:follower_id"`
	FollowingId int64 `gorm:"column:following_id"`
	IsDel       int8  `gorm:"column:is_del"`
}

func (Relation) TableName() string {
	return "dy_relation"
}
