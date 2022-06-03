package models

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"column:user_id"`
	Name          string `json:"name,omitempty" gorm:"column:name"`
	PassWord      string `gorm:"column:password" json:"password"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:follow_count"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:follower_count"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"column:is_follow"`
	Salt          string `json:"salt,omitempty" gorm:"column:salt"`
	Token         string `gorm:"-"`
}

func (User) TableName() string {
	return "dy_user"
}
