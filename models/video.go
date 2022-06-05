package models

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author" gorm:"-" sql:"-"` //不会入库，只做返回给前端的时候结构体
	UserId        int64  `json:"user_id" gorm:"user_id"`  //入库，实际存入的是userId
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title"`
}

func (Video) TableName() string {
	return "dy_video"
}
