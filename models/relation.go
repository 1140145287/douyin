package models

type DyRelation struct {
	Id          int64 `gorm:"column:relation_id"`
	FollowerId  int64 `gorm:"column:follower_id"`
	FollowingId int64 `gorm:"column:following_id"`
	IsDel       int64 `gorm:"column:is_del"`
}
